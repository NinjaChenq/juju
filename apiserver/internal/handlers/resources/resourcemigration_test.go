// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package resources

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/core/resource"
	domainresource "github.com/juju/juju/domain/resource"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/rpc/params"
)

const migrateResourcesPrefix = "/migrate/resources"

type resourcesUploadSuite struct {
	resourceServiceGetter *MockResourceServiceGetter
	resourceService       *MockResourceService

	mux *apiserverhttp.Mux
	srv *httptest.Server
}

var _ = gc.Suite(&resourcesUploadSuite{})

func (s *resourcesUploadSuite) SetUpTest(c *gc.C) {
	s.mux = apiserverhttp.NewMux()
	s.srv = httptest.NewServer(s.mux)
}

func (s *resourcesUploadSuite) TearDownTest(c *gc.C) {
	s.srv.Close()
}
func (s *resourcesUploadSuite) TestStub(c *gc.C) {
	c.Skip("This suite is missing tests for the following scenarios:\n" +
		"- Sending a POST req requires authorization via unit or application only.\n" +
		"- Rejects an unknown model with http.StatusNotFound.\n" +
		"- Test fails when model not importing.")
}

// TestServeMethodNotSupported ensures that the handler rejects HTTP methods
// other than POST with a 405 Method Not Allowed response.
func (s *resourcesUploadSuite) TestServeMethodNotSupported(c *gc.C) {
	// Arrange
	handler := NewResourceMigrationUploadHandler(
		nil,
		loggertesting.WrapCheckLog(c),
	)
	unsupportedMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	// This is a bit pathological, but we want to make sure that the handler
	// logic only actions on POST requests.
	for _, method := range unsupportedMethods {
		s.mux.AddHandler(method, migrateResourcesPrefix, handler)
		defer s.mux.RemoveHandler(method, migrateResourcesPrefix)
	}

	url := s.srv.URL + migrateResourcesPrefix
	for _, method := range unsupportedMethods {
		// Act
		request, err := http.NewRequest(method, url, nil)
		c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while building request. method: %s", method))
		response, err := http.DefaultClient.Do(request)
		c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request. method: %s", method))

		// Assert
		c.Check(response.StatusCode, gc.Equals, http.StatusMethodNotAllowed,
			gc.Commentf("(Assert) unexpected status code. method: %s", method))
	}
}

// TestServeUploadApplicationResourceNotFound verifies the handler's behavior
// when the application resource is not found.
func (s *resourcesUploadSuite) TestServeUploadApplicationResourceNotFound(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("", errors.NotFound)

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", nil)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusNotFound,
		gc.Commentf("(Assert) unexpected status code."))
}

// TestServeUploadApplicationStoreResourceError verifies error handling
// when storing a resource during an upload operation.
func (s *resourcesUploadSuite) TestServeUploadApplicationStoreResourceError(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), gomock.Any()).Return(errors.New("cannot store resource"))

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", nil)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusInternalServerError,
		gc.Commentf("(Assert) unexpected status code."))
}

// TestServeUploadApplicationGetResourceError validates the behavior
// when an error occurs while retrieving a resource.
func (s *resourcesUploadSuite) TestServeUploadApplicationGetResourceError(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), gomock.Any()).Return(nil)
	s.resourceService.EXPECT().GetResource(gomock.Any(), gomock.Any()).Return(resource.Resource{}, errors.New(
		"cannot get resource"))

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", nil)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusInternalServerError,
		gc.Commentf("(Assert) unexpected status code."))
}

// TestServeUploadApplicationWithPlaceholder tests the application's ability to
// handle uploading with placeholders correctly. It verifies that StoreResource is
// not called through not configuring related mock.
func (s *resourcesUploadSuite) TestServeUploadApplicationWithPlaceholder(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
	}

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusOK,
		gc.Commentf("(Assert) unexpected status code."))
}

// TestServeUploadApplication tests the HTTP endpoint for uploading application
// resources, ensuring correct handling and storage. It verifies correct values
// are passed as argument for calling the different underlying services.
func (s *resourcesUploadSuite) TestServeUploadApplication(c *gc.C) {
	// Arrange
	now := time.Now().Truncate(time.Second).UTC()
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), domainresource.StoreResourceArgs{
		ResourceUUID: "res-uuid",
		Reader:       http.NoBody,
	}).Return(nil)
	s.resourceService.EXPECT().GetResource(gomock.Any(), resource.UUID("res-uuid")).Return(resource.Resource{
		UUID:      "res-uuid",
		Timestamp: now,
	}, nil)

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	var obtained params.ResourceUploadResult
	c.Check(response.StatusCode, gc.Equals, http.StatusOK,
		gc.Commentf("(Assert) unexpected status code."))
	body, err := io.ReadAll(response.Body)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Assert) unexpected error while reading response body"))
	c.Assert(json.Unmarshal(body, &obtained), jc.ErrorIsNil,
		gc.Commentf("(Assert) unexpected error while unmarshalling response"))
	c.Check(obtained, gc.Equals, params.ResourceUploadResult{
		ID:        "res-uuid",
		Timestamp: now,
	})
}

// TestServeUploadApplicationRetrievedByUser tests that the RetrievedBy and
// RetrievedByType values are correctly determined for a user retriever.
func (s *resourcesUploadSuite) TestServeUploadApplicationRetrievedByUser(c *gc.C) {
	// Arrange
	now := time.Now().Truncate(time.Second).UTC()
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
		"user":        {"username"},
		"origin":      {"upload"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), domainresource.StoreResourceArgs{
		ResourceUUID:    "res-uuid",
		Reader:          http.NoBody,
		RetrievedByType: resource.User,
		RetrievedBy:     "username",
	}).Return(nil)
	s.resourceService.EXPECT().GetResource(gomock.Any(), resource.UUID("res-uuid")).Return(resource.Resource{
		UUID:      "res-uuid",
		Timestamp: now,
	}, nil)

	// Act
	_, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))
}

// TestServeUploadApplicationRetrievedByApplication tests that the RetrievedBy
// and RetrievedByType values are correctly determined for an application
// retriever.
func (s *resourcesUploadSuite) TestServeUploadApplicationRetrievedByApplication(c *gc.C) {
	// Arrange
	now := time.Now().Truncate(time.Second).UTC()
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
		"user":        {"app-name"},
		"origin":      {"store"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), domainresource.StoreResourceArgs{
		ResourceUUID:    "res-uuid",
		Reader:          http.NoBody,
		RetrievedByType: resource.Application,
		RetrievedBy:     "app-name",
	}).Return(nil)
	s.resourceService.EXPECT().GetResource(gomock.Any(), resource.UUID("res-uuid")).Return(resource.Resource{
		UUID:      "res-uuid",
		Timestamp: now,
	}, nil)

	// Act
	_, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))
}

// TestServeUploadApplicationRetrievedByUnit tests that the RetrievedBy and
// RetrievedByType values are correctly determined for a unit retriever.
func (s *resourcesUploadSuite) TestServeUploadApplicationRetrievedByUnit(c *gc.C) {
	// Arrange
	now := time.Now().Truncate(time.Second).UTC()
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":        {"resource-name"},
		"application": {"app-name"},
		"timestamp":   {"not-placeholder"},
		"user":        {"app-name/0"},
		"origin":      {"store"},
	}
	s.resourceService.EXPECT().GetResourceUUIDByApplicationAndResourceName(
		gomock.Any(),
		"app-name",
		"resource-name",
	).Return("res-uuid", nil)
	s.resourceService.EXPECT().StoreResource(gomock.Any(), domainresource.StoreResourceArgs{
		ResourceUUID:    "res-uuid",
		Reader:          http.NoBody,
		RetrievedByType: resource.Unit,
		RetrievedBy:     "app-name/0",
	}).Return(nil)
	s.resourceService.EXPECT().GetResource(gomock.Any(), resource.UUID("res-uuid")).Return(resource.Resource{
		UUID:      "res-uuid",
		Timestamp: now,
	}, nil)

	// Act
	_, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))
}

// TestServeUploadUnitWithPlaceholder tests the upload functionality for a unit
// with a placeholder in the resource upload service. It is basically the same
// test than the one with application, with one call to SetUnitResource.
func (s *resourcesUploadSuite) TestServeUploadUnitWithPlaceholder(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name": {"resource-name"},
		"unit": {"testunit-id"},
	}

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusOK,
		gc.Commentf("(Assert) unexpected status code."))
}

// TestServeUploadUnit tests the process of uploading a resource unit and
// verifies correct resource handling and response. It is basically the same
// test than the one with application, with one call to SetUnitResource.
func (s *resourcesUploadSuite) TestServeUploadUnit(c *gc.C) {
	// Arrange
	defer s.setupHandler(c).Finish()
	query := url.Values{
		"name":      {"resource-name"},
		"unit":      {"testunit-id"},
		"timestamp": {"not-placeholder"},
	}

	// Act
	response, err := http.Post(s.srv.URL+migrateResourcesPrefix+"?"+query.Encode(), "application/octet-stream", http.NoBody)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Act) unexpected error while executing request"))

	// Assert
	c.Check(response.StatusCode, gc.Equals, http.StatusOK,
		gc.Commentf("(Assert) unexpected status code."))
}

// setupHandler configures the resources migration upload HTTP handler, init
// mocks and registers it to the mux. It provides cleanup logic.
func (s *resourcesUploadSuite) setupHandler(c *gc.C) Finisher {
	finish := s.setupMocks(c).Finish
	s.expectResourceService()

	handler := NewResourceMigrationUploadHandler(
		s.resourceServiceGetter,
		loggertesting.WrapCheckLog(c),
	)

	err := s.mux.AddHandler("POST", migrateResourcesPrefix, handler)
	c.Assert(err, jc.ErrorIsNil, gc.Commentf("(Arrange) unexpected error while adding handler"))

	return &finisherWrapper{
		finish: func() {
			s.mux.RemoveHandler("POST", migrateResourcesPrefix)
			finish()
		},
	}
}

// expectResourceService prepare mocks for resource service
func (s *resourcesUploadSuite) expectResourceService() {
	s.resourceServiceGetter.EXPECT().Resource(gomock.Any()).Return(s.resourceService, nil)
}

// setupMocks initializes mock services and returns a gomock.Controller
// for managing mock lifecycle.
func (s *resourcesUploadSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.resourceServiceGetter = NewMockResourceServiceGetter(ctrl)
	s.resourceService = NewMockResourceService(ctrl)

	return ctrl
}

type Finisher interface {
	Finish()
}

type finisherWrapper struct {
	finish func()
}

func (f *finisherWrapper) Finish() {
	f.finish()
}
