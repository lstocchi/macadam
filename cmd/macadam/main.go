package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/crc-org/macadam/pkg/cmdline"

	"github.com/containers/common/pkg/config"
	ldefine "github.com/containers/podman/v5/libpod/define"

	"github.com/containers/podman/v5/pkg/machine"
	"github.com/containers/podman/v5/pkg/machine/define"
	"github.com/containers/podman/v5/pkg/machine/env"
	provider2 "github.com/containers/podman/v5/pkg/machine/provider"
	"github.com/containers/podman/v5/pkg/machine/shim"
	"github.com/containers/podman/v5/pkg/machine/vmconfigs"

	_ "github.com/spf13/cobra"
)

var (
	initOpts           = define.InitOptions{}
	defaultMachineName = machine.DefaultMachineName
)

type PodmanMachine struct {
	provider vmconfigs.VMProvider
	config   *vmconfigs.MachineConfig
}

func startMachine(m *PodmanMachine) error {

	machineName := m.config.Name
	dirs, err := env.GetMachineDirs(m.provider.VMType())
	if err != nil {
		return err
	}
	/*
		mc, err := vmconfigs.LoadMachineByName(machineName, dirs)
		if err != nil {
			return err
		}
	*/

	fmt.Printf("Starting machine %q\n", machineName)

	startOpts := machine.StartOptions{
		NoInfo: false,
		Quiet:  false,
	}
	if err := shim.Start(m.config, m.provider, dirs, startOpts); err != nil {
		return err
	}
	fmt.Printf("Machine %q started successfully\n", machineName)
	//newMachineEvent(events.Start, events.Event{Name: vmName})
	return nil
}

func initMachine(initOpts define.InitOptions) (*PodmanMachine, error) {
	machine := PodmanMachine{}
	provider, err := provider2.Get()
	if err != nil {
		return nil, err
	}
	machine.provider = provider

	// The vmtype names need to be reserved and cannot be used for podman machine names
	if _, err := define.ParseVMType(initOpts.Name, define.UnknownVirt); err == nil {
		return nil, fmt.Errorf("cannot use %q for a machine name", initOpts.Name)
	}

	if !ldefine.NameRegex.MatchString(initOpts.Username) {
		return nil, fmt.Errorf("invalid username %q: %w", initOpts.Username, ldefine.RegexError)
	}

	// Check if machine already exists
	vmConfig, exists, err := shim.VMExists(initOpts.Name, []vmconfigs.VMProvider{provider})
	if err != nil {
		return nil, err
	}
	machine.config = vmConfig

	// machine exists, return error
	if exists {
		return &machine, fmt.Errorf("%s: %w", initOpts.Name, define.ErrVMAlreadyExists)
	}

	/*
		// check if a system connection already exists
		cons, err := registry.PodmanConfig().ContainersConfDefaultsRO.GetAllConnections()
		if err != nil {
			return err
		}
		for _, con := range cons {
			if con.ReadWrite {
				for _, connection := range []string{initOpts.Name, fmt.Sprintf("%s-root", initOpts.Name)} {
					if con.Name == connection {
						return fmt.Errorf("system connection %q already exists. consider a different machine name or remove the connection with `podman system connection rm`", connection)
					}
				}
			}
		}
	*/

	for idx, vol := range initOpts.Volumes {
		initOpts.Volumes[idx] = os.ExpandEnv(vol)
	}

	// TODO need to work this back in
	// if finished, err := vm.Init(initOpts); err != nil || !finished {
	// 	// Finished = true,  err  = nil  -  Success! Log a message with further instructions
	// 	// Finished = false, err  = nil  -  The installation is partially complete and podman should
	// 	//                                  exit gracefully with no error and no success message.
	// 	//                                  Examples:
	// 	//                                  - a user has chosen to perform their own reboot
	// 	//                                  - reexec for limited admin operations, returning to parent
	// 	// Finished = *,     err != nil  -  Exit with an error message
	// 	return err
	// }

	err = shim.Init(initOpts, provider)
	if err != nil {
		return nil, err
	}

	/*
		newMachineEvent(events.Init, events.Event{Name: initOpts.Name})
	*/
	fmt.Println("Machine init complete")

	vmConfig, _, err = shim.VMExists(initOpts.Name, []vmconfigs.VMProvider{provider})
	if err != nil {
		return nil, err
	}
	machine.config = vmConfig

	/*
		now := false
		if now {
			return startMachine(initOpts.Name, provider)
		}
	*/
	extra := ""

	if initOpts.Name != defaultMachineName {
		extra = " " + initOpts.Name
	}
	fmt.Printf("To start your machine run:\n\n\tpodman machine start%s\n\n", extra)
	return &machine, err
}

func main() {
	slog.Info(fmt.Sprintf("macadam version %s", cmdline.Version()))

	defaultConfig, err := config.New(&config.Options{
		SetDefault: true, // This makes sure that following calls to config.Default() return this config
	})
	if err != nil {
		os.Exit(1)
	}

	// defaults from cmd/podman/machine/init.go
	initOpts.Name = defaultMachineName

	initOpts.CPUS = defaultConfig.Machine.CPUs
	initOpts.DiskSize = defaultConfig.Machine.DiskSize
	initOpts.Memory = defaultConfig.Machine.Memory
	defaultTz := defaultConfig.TZ()
	if len(defaultTz) < 1 {
		defaultTz = "local"
	}
	initOpts.TimeZone = defaultTz
	initOpts.ReExec = false
	initOpts.Username = defaultConfig.Machine.User
	initOpts.Image = defaultConfig.Machine.Image
	initOpts.Volumes = defaultConfig.Machine.Volumes.Get()
	initOpts.USBs = []string{}
	initOpts.VolumeDriver = ""
	initOpts.IgnitionPath = ""
	initOpts.Rootful = false
	userModeNetworking := false
	initOpts.UserModeNetworking = &userModeNetworking
	// user-mode networking

	machine, err := initMachine(initOpts)
	if err != nil && !errors.Is(err, define.ErrVMAlreadyExists) {
		slog.Error(err.Error())
	}
	if err := startMachine(machine); err != nil {
		slog.Error(err.Error())
	}
	/*
		if err != nil || errors.Is(err, define.ErrVMAlreadyExists)
		{
	*/
}
