// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"fmt"
	"strings"

	"launchpad.net/gnuflag"

	"launchpad.net/juju-core/cmd"
	"launchpad.net/juju-core/constraints"
	"launchpad.net/juju-core/environs/manual"
	"launchpad.net/juju-core/instance"
	"launchpad.net/juju-core/juju"
	"launchpad.net/juju-core/log"
	"launchpad.net/juju-core/names"
	"launchpad.net/juju-core/state"
)

// sshHostPrefix is the prefix for a machine to be "manually provisioned".
const sshHostPrefix = "ssh:"

var addMachineDoc = `
Machines are created in a clean state and ready to have units deployed.

This command also supports configuring existing machines via SSH. The
target machine must be able to communicate with the API servers, and
be able to access the environment storage.`[1:]

// AddMachineCommand starts a new machine and registers it in the environment.
type AddMachineCommand struct {
	cmd.EnvCommandBase
	// If specified, use this series, else use the environment default-series
	Series string
	// If specified, these constraints are merged with those already in the environment.
	Constraints   constraints.Value
	MachineId     string
	ContainerType instance.ContainerType
	SSHHost       string
}

func (c *AddMachineCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "add-machine",
		Args:    "[<container>:machine | <container> | ssh:[user@]host]",
		Purpose: "start a new, empty machine and optionally a container, or add a container to a machine",
		Doc:     addMachineDoc,
	}
}

func (c *AddMachineCommand) SetFlags(f *gnuflag.FlagSet) {
	c.EnvCommandBase.SetFlags(f)
	f.StringVar(&c.Series, "series", "", "the charm series")
	f.Var(constraints.ConstraintsValue{&c.Constraints}, "constraints", "additional machine constraints")
}

func (c *AddMachineCommand) Init(args []string) error {
	if c.Constraints.Container != nil {
		return fmt.Errorf("container constraint %q not allowed when adding a machine", *c.Constraints.Container)
	}
	containerSpec, err := cmd.ZeroOrOneArgs(args)
	if err != nil {
		return err
	}
	if containerSpec == "" {
		return nil
	}
	if strings.HasPrefix(containerSpec, sshHostPrefix) {
		c.SSHHost = containerSpec[len(sshHostPrefix):]
	} else {
		// container arg can either be 'type:machine' or 'type'
		if c.ContainerType, err = instance.ParseSupportedContainerType(containerSpec); err != nil {
			if names.IsMachine(containerSpec) || !cmd.IsMachineOrNewContainer(containerSpec) {
				return fmt.Errorf("malformed container argument %q", containerSpec)
			}
			sep := strings.Index(containerSpec, ":")
			c.MachineId = containerSpec[sep+1:]
			c.ContainerType, err = instance.ParseSupportedContainerType(containerSpec[:sep])
		}
	}
	return err
}

func (c *AddMachineCommand) Run(_ *cmd.Context) error {
	conn, err := juju.NewConnFromName(c.EnvName)
	if err != nil {
		return err
	}
	defer conn.Close()

	if c.SSHHost != "" {
		args := manual.ProvisionMachineArgs{
			Host:        c.SSHHost,
			Env:         conn.Environ,
			State:       conn.State,
			Constraints: c.Constraints,
		}
		_, err = manual.ProvisionMachine(args)
		return err
	}

	series := c.Series
	if series == "" {
		conf, err := conn.State.EnvironConfig()
		if err != nil {
			return err
		}
		series = conf.DefaultSeries()
	}
	params := state.AddMachineParams{
		ParentId:      c.MachineId,
		ContainerType: c.ContainerType,
		Series:        series,
		Constraints:   c.Constraints,
		Jobs:          []state.MachineJob{state.JobHostUnits},
	}
	m, err := conn.State.AddMachineWithConstraints(&params)
	if err == nil {
		if c.ContainerType == "" {
			log.Infof("created machine %v", m)
		} else {
			log.Infof("created %q container on machine %v", c.ContainerType, m)
		}
	}
	return err
}
