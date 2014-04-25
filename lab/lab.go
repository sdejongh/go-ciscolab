package lab

import (
	"ciscolab/devices"
	"ciscolab/utilities"
	"sort"
)

// Lab structure, contains slices of devices and comservers
type Lab struct {
	labComServers []*devices.ComServer
	labDevices    []*devices.LabDevice
}

// Returns number of comservers already added
func (l Lab) GetComServersCount() int { return len(l.labComServers) }

// Returns number of devices already added
func (l Lab) GetDevicesCount() int { return len(l.labDevices) }

// Adds *ComServer's to the slice
func (l *Lab) AddComServer(d ...*devices.ComServer) {
	l.labComServers = append(l.labComServers, d...)
}

// Adds *Device's to the slice
func (l *Lab) AddDevice(d ...*devices.LabDevice) {
	l.labDevices = append(l.labDevices, d...)
}

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

// Lab type constructor
func NewLab(maxComServers int16, maxDevices int16) *Lab {
	comservers := make([]*devices.ComServer, 0, maxComServers)
	devices := make([]*devices.LabDevice, 0, maxDevices)
	return &Lab{labComServers: comservers, labDevices: devices}
}
