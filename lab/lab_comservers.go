/*
LAB package
type/functions/... related to the ComServer management
*/

package lab

import (
	"ciscolab/devices"
	"ciscolab/utilities"
	"sort"
)

// Adds *ComServer's to the slice
func (l *Lab) AddComServer(d ...*devices.ComServer) { l.labComServers = append(l.labComServers, d...) }

// Returns number of comservers added to the lab
func (l Lab) GetComServersCount() int { return len(l.labComServers) }

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

// Checks whether or not the ComServer exists (lookup by name)
func (l *Lab) IsComServerInLab(name string) bool {
	for _, v := range l.ListComServersByName(false) {
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
