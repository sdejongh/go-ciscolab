package lab

import (
	"ciscolab/devices"
	"ciscolab/utilities"
	"sort"
)

// Lab structure, contains slices for devices and comservers
type Lab struct {
	labComServers []*devices.ComServer
	labDevices    []*devices.LabDevice
}

// Lab type constructor
func NewLab(maxComServers int16, maxDevices int16) *Lab {
	comservers := make([]*devices.ComServer, 0, maxComServers)
	devices := make([]*devices.LabDevice, 0, maxDevices)
	return &Lab{labComServers: comservers, labDevices: devices}
}

// Returns number of comservers added to the lab
func (l Lab) GetComServersCount() int { return len(l.labComServers) }

// Returns number of devices added to the lab
func (l Lab) GetDevicesCount() int { return len(l.labDevices) }

// Adds *ComServer's to the slice
func (l *Lab) AddComServer(d ...*devices.ComServer) { l.labComServers = append(l.labComServers, d...) }

// Adds *Device's to the slice
func (l *Lab) AddDevice(d ...*devices.LabDevice) { l.labDevices = append(l.labDevices, d...) }

// Returns a slice containing ComServers name, sorted or not.
func (l Lab) ListComServersByName(sorted bool) []string {
	nComServers := l.GetComServersCount()
	list := make([]string, nComServers)
	for i, v := range l.labComServers {
		list[i] = v.GetName()
	}
	if sorted {
		sort.Sort(utilities.StringsByName(list))
	}
	return list
}

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

// Checks whether or not the ComServer exists (lookup by name)
func (l *Lab) IsComServerInLab(name string) bool {
	for _, v := range l.ListComServersByName(false) {
		if v == name {
			return true
		}
	}
	return false
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

// Returns the pointer corresponding to the given comserver name
func (l *Lab) GetComServerByName(name string) *devices.ComServer {
	if l.IsComServerInLab(name) {
		for _, v := range l.labComServers {
			if v.GetName() == name {
				return v
			}
		}
	}
	return nil
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
