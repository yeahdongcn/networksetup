//go:build darwin
// +build darwin

package vlan

import (
	"net"
	"time"
)

func SettleAddresses(ifName string) error {
	ifs, err := net.Interfaces()
	if err != nil {
		return err
	} else {
		for _, iface := range ifs {
			if iface.Name == ifName {
				for {
					addrs, err := iface.Addrs()
					if err != nil {
						return err
					}
					ipv4Ready := false
					ipv6Ready := false
					for _, addr := range addrs {
						if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
							if ip := ipnet.IP.To4(); ip != nil {
								ipv4Ready = true
							} else if ipnet.IP.To16() != nil {
								ipv6Ready = true
							}
						}
					}
					if ipv4Ready && ipv6Ready {
						return nil
					}
					// Sleep for a while before checking again
					time.Sleep(time.Second)
				}
			}
		}
	}
	return nil
}
