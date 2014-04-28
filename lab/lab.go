/*
LAB package
type/functions/... related to general management
*/

package lab

import (
	"ciscolab/devices"
	"ciscolab/utilities"
	"reflect"
	"sort"
)

// Lab structure, contains slices for devices and comservers
type Lab struct {
	devices map[string]devices.Devicer
}

// Lab type constructor
func NewLab() *Lab {
	devices := make(map[string]devices.Devicer)
	return &Lab{devices: devices}
}

// Add ComServer's and/or LabDevice's to the lab
func (l *Lab) AddDevice(devicers ...devices.Devicer) {
	for _, device := range devicers {
		l.devices[device.GetName()] = device
	}
}

// Remove ComServer or LabDevice from Lab (by name)
func (l *Lab) RemoveDevice(name string) {
	if l.IsDeviceInLab(name) {
		delete(l.devices, name)
	}
}

// Return all the ComServer's names sorted, or not...
func (l Lab) GetComServerNames(sorted bool) []string {
	names := make([]string, len(l.devices))
	for _, d := range l.devices {
		v := reflect.TypeOf(d).Elem()
		dummyComServer := devices.ComServer{}
		if v == reflect.TypeOf(dummyComServer) {
			names = append(names, d.GetName())
		}
	}
	if sorted {
		sort.Sort(utilities.StringsByName(names))
	}
	return names
}

// Return all the LabDevice's names sorted, or not...
func (l Lab) GetLabDeviceNames(sorted bool) []string {
	names := make([]string, len(l.devices))
	for _, d := range l.devices {
		v := reflect.TypeOf(d).Elem()
		dummyLabDevice := devices.LabDevice{}
		if v == reflect.TypeOf(dummyLabDevice) {
			names = append(names, d.GetName())
		}
	}
	if sorted {
		sort.Sort(utilities.StringsByName(names))
	}
	return names
}

// Checks whether the device is in the lab or not
func (l Lab) IsDeviceInLab(name string) bool {
	_, exists := l.devices[name]
	return exists
}

// Returns the ptr for the given device name
func (l Lab) GetDeviceByName(name string) devices.Devicer {
	device, _ := l.devices[name]
	return device
}
