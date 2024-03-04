//go:build darwin
// +build darwin

package vlan

import (
	"os/exec"
	"strconv"
)

func Delete(name, parent string, tag int) error {
	cmd := exec.Command(networksetupCommand)
	cmd.Args = append(cmd.Args, "-deleteVLAN")
	cmd.Args = append(cmd.Args, name)
	cmd.Args = append(cmd.Args, parent)
	cmd.Args = append(cmd.Args, strconv.Itoa(tag))
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
