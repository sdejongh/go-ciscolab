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
	labComServers []*devices.ComServer
	labDevices    []*devices.LabDevice
}

// Lab type constructor
func NewLab(maxComServers int16, maxDevices int16) *Lab {
	comservers := make([]*devices.ComServer, 0, maxComServers)
	devices := make([]*devices.LabDevice, 0, maxDevices)
	return &Lab{labComServers: comservers, labDevices: devices}
}
