//go:build darwin
// +build darwin

package vlan

func Get(name string) *Device {
	devices, err := List()
	if err != nil {
		return nil
	}
	for _, device := range devices {
		if device.Name == name {
			return &device
		}
	}
	return nil
}
