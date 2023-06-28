// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package leaseexpiry

import (
	"context"
	"time"

	"github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/worker/v3"
	"gopkg.in/tomb.v2"

	"github.com/juju/juju/core/lease"
)

// Config encapsulates the configuration options for
// instantiating a new lease expiry worker.
type Config struct {
	Clock  clock.Clock
	Logger Logger
	Store  lease.ExpiryStore
}

// Validate checks whether the worker configuration settings are valid.
func (cfg Config) Validate() error {
	if cfg.Clock == nil {
		return errors.NotValidf("nil Clock")
	}
	if cfg.Logger == nil {
		return errors.NotValidf("nil Logger")
	}
	if cfg.Store == nil {
		return errors.NotValidf("nil Store")
	}

	return nil
}

type expiryWorker struct {
	tomb tomb.Tomb

	clock  clock.Clock
	logger Logger
	store  lease.ExpiryStore
}

// NewWorker returns a worker that periodically deletes
// expired leases from the controller database.
func NewWorker(cfg Config) (worker.Worker, error) {
	var err error

	if err = cfg.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	w := &expiryWorker{
		clock:  cfg.Clock,
		logger: cfg.Logger,
		store:  cfg.Store,
	}

	w.tomb.Go(w.loop)
	return w, nil

}

func (w *expiryWorker) loop() error {
	timer := w.clock.NewTimer(time.Second)
	defer timer.Stop()

	// We pass this context to every database method that accepts one.
	// It is cancelled by killing the tomb, which prevents shutdown
	// being blocked by such calls.
	ctx := w.tomb.Context(context.Background())

	for {
		select {
		case <-w.tomb.Dying():
			return tomb.ErrDying
		case <-timer.Chan():
			if err := w.store.ExpireLeases(ctx); err != nil {
				return errors.Trace(err)
			}
			timer.Reset(time.Second)
		}
	}
}

// Kill is part of the worker.Worker interface.
func (w *expiryWorker) Kill() {
	w.tomb.Kill(nil)
}

// Wait is part of the worker.Worker interface.
func (w *expiryWorker) Wait() error {
	return w.tomb.Wait()
}
