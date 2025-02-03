// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasadmission

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/juju/errors"
	"github.com/juju/worker/v4/catacomb"

	"github.com/juju/juju/core/logger"
)

type Mux interface {
	AddHandler(string, string, http.Handler) error
	RemoveHandler(string, string)
}

// Kubernetes controller responsible
type Controller struct {
	catacomb catacomb.Catacomb
	logger   logger.Logger
}

func AdmissionPathForModel(modelUUID string) string {
	return fmt.Sprintf("/k8s/admission/%s", url.PathEscape(modelUUID))
}

func NewController(
	logger logger.Logger,
	mux Mux,
	path string,
	legacyLabels bool,
	admissionCreator AdmissionCreator,
	rbacMapper RBACMapper) (*Controller, error) {

	c := &Controller{
		logger: logger,
	}

	if err := catacomb.Invoke(catacomb.Plan{
		Site: &c.catacomb,
		Work: c.makeLoop(admissionCreator,
			admissionHandler(logger, rbacMapper, legacyLabels),
			logger, mux, path),
	}); err != nil {
		return c, errors.Trace(err)
	}

	return c, nil
}

func (c *Controller) Kill() {
	c.catacomb.Kill(nil)
}

func (c *Controller) makeLoop(
	admissionCreator AdmissionCreator,
	handler http.Handler,
	logger logger.Logger,
	mux Mux,
	path string) func() error {

	return func() error {
		logger.Debugf(context.TODO(), "installing caas admission handler at %s", path)
		if err := mux.AddHandler(http.MethodPost, path, handler); err != nil {
			return errors.Trace(err)
		}
		defer mux.RemoveHandler(http.MethodPost, path)

		logger.Infof(context.TODO(), "ensuring model k8s webhook configurations")
		admissionCleanup, err := admissionCreator.EnsureMutatingWebhookConfiguration(context.TODO())
		if err != nil {
			return errors.Trace(err)
		}
		defer admissionCleanup()

		select {
		case <-c.catacomb.Dying():
			return c.catacomb.ErrDying()
		}
	}
}

func (c *Controller) Wait() error {
	return c.catacomb.Wait()
}
