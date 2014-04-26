/*
LAB package
type/functions/... related to general management
*/

package lab

import (
	"ciscolab/devices"
)

// Lab structure, contains slices for devices and comservers
type Lab struct {
	labComServers map[string]*devices.ComServer
	labDevices    map[string]*devices.LabDevice
}

// Lab type constructor
func NewLab() *Lab {
	comservers := make(map[string]*devices.ComServer)
	devices := make(map[string]*devices.LabDevice)
	return &Lab{labComServers: comservers, labDevices: devices}
}
