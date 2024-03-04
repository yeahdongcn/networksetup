package main

import (
	"log"

	"github.com/yeahdongcn/networksetup/vlan"
)

func main() {
	pds, err := vlan.ListSupportedParentDevices()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Parent devices that support VLAN: %#v\n", pds)
	}

	if len(pds) == 0 {
		log.Fatal("No device supports VLAN")
	}

	if err := vlan.Create("vlan1", pds[0].IfName, 1); err != nil {
		log.Fatal(err)
	} else {
		log.Println("VLAN created")
	}

	dev := vlan.Get("vlan1")
	if dev == nil {
		log.Fatal("VLAN device not found")
	} else {
		log.Printf("VLAN device: %#v\n", dev)
	}

	if err := vlan.SettleAddresses(dev.IfName); err != nil {
		log.Fatal(err)
	} else {
		log.Println("VLAN addresses settled")
	}

	devices, err := vlan.List()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("VLAN devices: %#v\n", devices)
	}

	if len(devices) == 0 {
		log.Println("No VLAN device found")
	}

	err = vlan.Delete("vlan1", pds[0].IfName, 1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("VLAN deleted")
	}
}
