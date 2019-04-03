// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package instancemutater

import (
	"fmt"

	"github.com/juju/juju/core/lxdprofile"

	"github.com/juju/errors"
	"gopkg.in/juju/names.v2"

	"github.com/juju/juju/api/base"
	apiwatcher "github.com/juju/juju/api/watcher"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/core/watcher"
)

//go:generate mockgen -package mocks -destination mocks/caller_mock.go github.com/juju/juju/api/base APICaller,FacadeCaller
//go:generate mockgen -package mocks -destination mocks/machinemutater_mock.go github.com/juju/juju/api/instancemutater MutaterMachine
type MutaterMachine interface {

	// InstanceId returns the provider specific instance id for this machine
	InstanceId() (string, error)

	// CharmProfilingInfo returns info to update lxd profiles on the machine
	CharmProfilingInfo() (*UnitProfileInfo, error)

	// SetCharmProfiles records the given slice of charm profile names.
	SetCharmProfiles([]string) error

	// SetUpgradeCharmProfileComplete records the result of updating
	// the machine's charm profile(s), for the given unit.
	SetUpgradeCharmProfileComplete(unitName string, message string) error

	// Tag returns the current machine tag
	Tag() names.MachineTag

	// RemoveUpgradeCharmProfileData completely removes the instance charm
	// profile data for a machine and the given unit, even if the machine
	// is dead.
	RemoveUpgradeCharmProfileData(string) error

	// WatchUnits returns a watcher.StringsWatcher for watching units of a given
	// machine.
	WatchUnits() (watcher.StringsWatcher, error)

	// WatchApplicationLXDProfiles returns a NotifyWatcher, notifies when the
	// following changes happen:
	//  - application charm URL changes and there is a lxd profile
	//  - unit is add or removed and there is a lxd profile
	WatchApplicationLXDProfiles() (watcher.NotifyWatcher, error)

	// SetModificationStatus sets the provider specific modification status
	// for a machine. Allowing the propagation of status messages to the
	// operator.
	SetModificationStatus(status status.Status, info string, data map[string]interface{}) error
}

// Machine represents a juju machine as seen by an instancemutater
// worker.
type Machine struct {
	facade base.FacadeCaller

	tag  names.MachineTag
	life params.Life
}

// InstanceId implements MutaterMachine.InstanceId.
func (m *Machine) InstanceId() (string, error) {
	var results params.StringResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: m.tag.String()}},
	}
	err := m.facade.FacadeCall("InstanceId", args, &results)
	if err != nil {
		return "", err
	}
	if len(results.Results) != 1 {
		return "", errors.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return "", result.Error
	}
	return result.Result, nil
}

// SetCharmProfiles implements MutaterMachine.SetCharmProfiles.
func (m *Machine) SetCharmProfiles([]string) error {
	return nil
}

// SetUpgradeCharmProfileComplete implements MutaterMachine.SetUpgradeCharmProfileComplete.
func (m *Machine) SetUpgradeCharmProfileComplete(unitName string, message string) error {
	var results params.ErrorResults
	args := params.SetProfileUpgradeCompleteArgs{
		Args: []params.SetProfileUpgradeCompleteArg{
			{
				Entity:   params.Entity{Tag: m.tag.String()},
				UnitName: unitName,
				Message:  message,
			},
		},
	}
	err := m.facade.FacadeCall("SetUpgradeCharmProfileComplete", args, &results)
	if err != nil {
		return err
	}
	if len(results.Results) != 1 {
		return errors.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Tag implements MutaterMachine.Tag.
func (m *Machine) Tag() names.MachineTag {
	return m.tag
}

// WatchUnits implements MutaterMachine.WatchUnits.
func (m *Machine) WatchUnits() (watcher.StringsWatcher, error) {
	var results params.StringsWatchResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: m.tag.String()}},
	}
	err := m.facade.FacadeCall("WatchUnits", args, &results)
	if err != nil {
		return nil, err
	}
	if len(results.Results) != 1 {
		return nil, fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return nil, result.Error
	}
	w := apiwatcher.NewStringsWatcher(m.facade.RawAPICaller(), result)
	return w, nil
}

// WatchApplicationLXDProfiles implements MutaterMachine.WatchApplicationLXDProfiles.
func (m *Machine) WatchApplicationLXDProfiles() (watcher.NotifyWatcher, error) {
	var results params.NotifyWatchResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: m.tag.String()}},
	}
	err := m.facade.FacadeCall("WatchApplicationLXDProfiles", args, &results)
	if err != nil {
		return nil, err
	}
	if len(results.Results) != 1 {
		return nil, fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return nil, result.Error
	}
	return apiwatcher.NewNotifyWatcher(m.facade.RawAPICaller(), result), nil
}

type UnitProfileInfo struct {
	ModelName       string
	InstanceId      instance.Id
	ProfileChanges  []lxdprofile.Profile
	CurrentProfiles []string
}

// CharmProfilingInfo implements MutaterMachine.CharmProfilingInfo.
func (m *Machine) CharmProfilingInfo() (*UnitProfileInfo, error) {
	var result params.CharmProfilingInfoResult
	args := params.Entity{Tag: m.tag.String()}
	err := m.facade.FacadeCall("CharmProfilingInfo", args, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.Trace(result.Error)
	}
	returnResult := &UnitProfileInfo{
		InstanceId:      result.InstanceId,
		ModelName:       result.ModelName,
		CurrentProfiles: result.CurrentProfiles,
	}
	profileChanges := make([]lxdprofile.Profile, len(result.ProfileChanges))
	for i, change := range result.ProfileChanges {
		profileChanges[i] = lxdprofile.Profile{
			Config:      change.Profile.Config,
			Description: change.Profile.Description,
			Devices:     change.Profile.Devices,
		}
		if change.Error != nil {
			return nil, change.Error
		}
	}
	returnResult.ProfileChanges = profileChanges
	return returnResult, nil
}

// RemoveUpgradeCharmProfileData implements MutaterMachine.RemoveUpgradeCharmProfileData.
func (m *Machine) RemoveUpgradeCharmProfileData(string) error {
	return nil
}

// SetModificationStatus implements MutaterMachine.SetModificationStatus.
func (m *Machine) SetModificationStatus(status status.Status, info string, data map[string]interface{}) error {
	var result params.ErrorResults
	args := params.SetStatus{
		Entities: []params.EntityStatusArgs{
			{Tag: m.tag.String(), Status: status.String(), Info: info, Data: data},
		},
	}
	err := m.facade.FacadeCall("SetModificationStatus", args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}
