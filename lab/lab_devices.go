/*
LAB package
type/functions/... related to the LabDevice management
*/

package lab

import (
	"ciscolab/devices"
	"ciscolab/utilities"
	"sort"
)

// Adds *Device's to the slice
func (l *Lab) AddDevice(d ...*devices.LabDevice) { l.labDevices = append(l.labDevices, d...) }

// Returns number of devices added to the lab
func (l Lab) GetDevicesCount() int { return len(l.labDevices) }

// Returns a slice containing Devices name, sorted or not
func (l Lab) ListDevicesByName(sorted bool) []string {
	nDevices := l.GetDevicesCount()
	list := make([]string, nDevices)
	for i, v := range l.labDevices {
		list[i] = v.GetName()
	}
	if sorted {
		sort.Sort(utilities.StringsByName(list))
	}
	return list
}

// Checks whether or not the device exists (lookup by name)
func (l *Lab) IsDeviceInLab(name string) bool {
	for _, v := range l.ListDevicesByName(false) {
		if v == name {
			return true
		}
	}
	return false
}

// Returns the pointer corresponding to the given device name
func (l *Lab) GetDeviceByName(name string) *devices.LabDevice {
	if l.IsDeviceInLab(name) {
		for _, v := range l.labDevices {
			if v.GetName() == name {
				return v
			}
		}
	}
	return nil
}
