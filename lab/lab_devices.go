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

// Adds *LabDevice to the map
func (l *Lab) AddDevice(devices ...*devices.LabDevice) {
	for _, device := range devices {
		l.labDevices[device.GetName()] = device
	}
}

// Removes *LabDevice from map
func (l *Lab) RemoveDevice(name string) {
	if l.IsDeviceInLab(name) {
		delete(l.labDevices, name)
	}
}

// Returns number of devices added to the lab
func (l Lab) GetDevicesCount() int { return len(l.labDevices) }

// Returns a slice containing Devices name, sorted or not.
func (l Lab) ListDevicesByName(sorted bool) []string {
	names := make([]string, len(l.labDevices))
	for i, _ := range l.labDevices {
		names = append(names, i)
	}
	if sorted {
		sort.Sort(utilities.StringsByName(names))
	}
	return names
}

// Checks whether or not the Device exists (lookup by name)
func (l *Lab) IsDeviceInLab(name string) bool {
	_, exists := l.labDevices[name]
	return exists
}

// Returns the pointer corresponding to the given device name
func (l *Lab) GetDeviceByName(name string) *devices.LabDevice {
	ptr, _ := l.labDevices[name]
	return ptr
}
