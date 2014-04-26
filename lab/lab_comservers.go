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

// Adds *ComServer to the map
func (l *Lab) AddComServer(comservers ...*devices.ComServer) {
	for _, comserver := range comservers {
		l.labComServers[comserver.GetName()] = comserver
	}
}

// Removes *ComServer from map
func (l *Lab) RemoveComServer(name string) {
	if l.IsComServerInLab(name) {
		delete(l.labComServers, name)
	}
}

// Returns number of comservers added to the lab
func (l Lab) GetComServersCount() int { return len(l.labComServers) }

// Returns a slice containing ComServers name, sorted or not.
func (l Lab) ListComServersByName(sorted bool) []string {
	names := make([]string, len(l.labComServers))
	for i, _ := range l.labComServers {
		names = append(names, i)
	}
	if sorted {
		sort.Sort(utilities.StringsByName(names))
	}
	return names
}

// Checks whether or not the ComServer exists (lookup by name)
func (l *Lab) IsComServerInLab(name string) bool {
	_, exists := l.labComServers[name]
	return exists
}

// Returns the pointer corresponding to the given comserver name
func (l *Lab) GetComServerByName(name string) *devices.ComServer {
	ptr, _ := l.labComServers[name]
	return ptr
}
