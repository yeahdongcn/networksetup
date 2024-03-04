//go:build darwin
// +build darwin

package vlan

import (
	"bufio"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type ParentDevice struct {
	IfName  string
	DevName string
}

type Device struct {
	Name         string
	ParentIfName string
	IfName       string
	Tag          int
}

func List() ([]Device, error) {
	var devices []Device

	cmd := exec.Command(networksetupCommand)
	cmd.Args = append(cmd.Args, "-listVLANs")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return devices, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ": ", 2)
		if len(parts) < 2 {
			continue
		}
		switch parts[0] {
		case "VLAN User Defined Name":
			devices = append(devices, Device{Name: parts[1]})
		case "Parent Device":
			devices[len(devices)-1].ParentIfName = parts[1]
		case "Device (\"Hardware\" Port)":
			devices[len(devices)-1].IfName = parts[1]
		case "Tag":
			tag, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			devices[len(devices)-1].Tag = tag
		}
	}

	return devices, nil
}

func ListSupportedParentDevices() ([]ParentDevice, error) {
	var pds []ParentDevice

	cmd := exec.Command(networksetupCommand)
	cmd.Args = append(cmd.Args, "-listdevicesthatsupportVLAN")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return pds, err
	}

	r := regexp.MustCompile(`(\w+)\s+\((.+)\)`)
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())
		if len(matches) < 3 {
			continue
		}
		pds = append(pds, ParentDevice{
			IfName:  matches[1],
			DevName: matches[2],
		})
	}

	return pds, nil
}
