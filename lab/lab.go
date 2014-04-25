package lab

import (
	"ciscolab/devices"
	"sort"
)

// Lab structure, contains slices of devices and comservers
type Lab struct {
	labComServers []*devices.ComServer
	labDevices    []*devices.LabDevice
}

// Type for sorted list of comservers by name
type LabComServersByName []string

// Returns number of comservers in slice (needed for sort.Sort)
func (n LabComServersByName) Len() int { return len(n) }

// Swaps two comservers in slice (needed for sort.Sort)
func (n LabComServersByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// checks which Comserver name is lesser than the other one (needed for sort.Sort)
func (n LabComServersByName) Less(i, j int) bool { return n[i] < n[j] }

// Adds *ComServer's to the slice
func (l *Lab) AddComServer(d ...*devices.ComServer) {
	l.labComServers = append(l.labComServers, d...)
}

// Returns a slice containing ComServers name
func (l Lab) ListComServersByName(sorted bool) []string {
	nComServers := len(l.labComServers)
	list := make([]string, nComServers)

	for i, v := range l.labComServers {
		list[i] = v.GetName()
	}
	if sorted {
		sort.Sort(LabComServersByName(list))
	}

	return list

}

// Lab type constructor
func NewLab(maxComServers int16, maxDevices int16) *Lab {
	comservers := make([]*devices.ComServer, 0, maxComServers)
	devices := make([]*devices.LabDevice, 0, maxDevices)
	return &Lab{labComServers: comservers, labDevices: devices}
}
