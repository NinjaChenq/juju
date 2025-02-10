// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package providertracker

import (
	"testing"

	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package providertracker -destination provider_mock_test.go github.com/juju/juju/core/providertracker ProviderFactory,Provider,NonTrackedProvider

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}
