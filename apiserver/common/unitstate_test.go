// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common_test

import (
	"context"

	"github.com/juju/names/v6"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/mocks"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/domain/unitstate"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/testing"
	"github.com/juju/juju/rpc/params"
)

type unitStateSuite struct {
	testing.BaseSuite

	unitTag1 names.UnitTag
	api      *common.UnitStateAPI

	controllerConfigGetter *mocks.MockControllerConfigService
	unitStateService       *mocks.MockUnitStateService
}

var _ = gc.Suite(&unitStateSuite{})

func (s *unitStateSuite) SetUpTest(c *gc.C) {
	s.unitTag1 = names.NewUnitTag("wordpress/0")
}

func (s *unitStateSuite) assertBackendApi(c *gc.C) *gomock.Controller {
	resources := common.NewResources()
	authorizer := apiservertesting.FakeAuthorizer{
		Tag: s.unitTag1,
	}

	ctrl := gomock.NewController(c)
	s.controllerConfigGetter = mocks.NewMockControllerConfigService(ctrl)
	s.unitStateService = mocks.NewMockUnitStateService(ctrl)

	unitAuthFunc := func() (common.AuthFunc, error) {
		return func(tag names.Tag) bool {
			if tag.Id() == s.unitTag1.Id() {
				return true
			}
			return false
		}, nil
	}

	s.api = common.NewUnitStateAPI(
		s.controllerConfigGetter,
		s.unitStateService,
		resources,
		authorizer,
		unitAuthFunc,
		loggertesting.WrapCheckLog(c),
	)
	return ctrl
}

func (s *unitStateSuite) expectGetState(name string) (map[string]string, string, map[int]string, string, string) {
	expCharmState := map[string]string{
		"foo.bar":  "baz",
		"payload$": "enc0d3d",
	}
	expUniterState := "testing"
	expRelationState := map[int]string{
		1: "one",
		2: "two",
	}
	expStorageState := "storage testing"
	expSecretState := "secret testing"

	uuid := "some-unit-uuid"
	s.unitStateService.EXPECT().GetUnitUUIDForName(gomock.Any(), name).Return(uuid, nil)
	s.unitStateService.EXPECT().GetState(gomock.Any(), uuid).Return(unitstate.RetrievedUnitState{
		CharmState:    expCharmState,
		UniterState:   expUniterState,
		RelationState: expRelationState,
		StorageState:  expStorageState,
		SecretState:   expSecretState,
	}, nil)

	return expCharmState, expUniterState, expRelationState, expStorageState, expSecretState
}

func (s *unitStateSuite) TestState(c *gc.C) {
	defer s.assertBackendApi(c).Finish()
	expCharmState, expUniterState, expRelationState, expStorageState, expSecretState := s.expectGetState("wordpress/0")

	args := params.Entities{
		Entities: []params.Entity{
			{Tag: "not-a-unit-tag"},
			{Tag: "unit-wordpress-0"},
			{Tag: "unit-mysql-0"}, // not accessible by current user
			{Tag: "unit-notfound-0"},
		},
	}
	result, err := s.api.State(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, params.UnitStateResults{
		Results: []params.UnitStateResult{
			{Error: &params.Error{Message: `"not-a-unit-tag" is not a valid tag`}},
			{
				Error:         nil,
				CharmState:    expCharmState,
				UniterState:   expUniterState,
				RelationState: expRelationState,
				StorageState:  expStorageState,
				SecretState:   expSecretState,
			},
			{Error: apiservertesting.ErrUnauthorized},
			{Error: apiservertesting.ErrUnauthorized},
		},
	})
}

func (s *unitStateSuite) TestSetStateUniterState(c *gc.C) {
	defer s.assertBackendApi(c).Finish()
	expUniterState := "testing"

	args := params.SetUnitStateArgs{
		Args: []params.SetUnitStateArg{
			{Tag: "not-a-unit-tag", UniterState: &expUniterState},
			{Tag: "unit-wordpress-0", UniterState: &expUniterState},
			{Tag: "unit-mysql-0", UniterState: &expUniterState}, // not accessible by current user
			{Tag: "unit-notfound-0", UniterState: &expUniterState},
		},
	}

	expectedState := unitstate.UnitState{
		Name:        "wordpress/0",
		UniterState: &expUniterState,
	}
	s.unitStateService.EXPECT().SetState(gomock.Any(), expectedState).Return(nil)

	result, err := s.api.SetState(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{
			{Error: &params.Error{Message: `"not-a-unit-tag" is not a valid tag`}},
			{Error: nil},
			{Error: apiservertesting.ErrUnauthorized},
			{Error: apiservertesting.ErrUnauthorized},
		},
	})
}
