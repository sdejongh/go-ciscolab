/*
LAB package
type/functions/... related to general management
*/

package lab

import (
	"ciscolab/config"
	"ciscolab/devices"
	"ciscolab/menu"
	"ciscolab/utilities"
	"sort"
	"strconv"
	"strings"
)

// Lab structure, contains slices for devices and comservers
type Lab struct {
	devices  map[string]devices.Devicer
	mainMenu menu.Menu
}

// Lab type constructor
func NewLab() *Lab {
	devices := make(map[string]devices.Devicer)
	return &Lab{devices: devices}
}

// Build lab from package "config"
func (l *Lab) BuildFromPackage() {
	for csName, csInfos := range config.ComServers {
		arrInfos := strings.Split(csInfos, ":")
		ipAddress := arrInfos[0]
		telnetPort, _ := strconv.Atoi(arrInfos[1])
		l.AddDevice(devices.NewComServer(csName, uint16(telnetPort), ipAddress))
	}
	for dName, dInfos := range config.LabDevices {
		arrInfos := strings.Split(dInfos, ":")
		comserver := l.GetDeviceByName(arrInfos[0])
		line, _ := strconv.Atoi(arrInfos[1])
		l.AddDevice(devices.NewLabDevice(dName, int16(line), comserver.(*devices.ComServer)))
	}
}

// Add ComServer's and/or LabDevice's to the lab
func (l *Lab) AddDevice(devicer devices.Devicer) {
	l.devices[devicer.GetName()] = devicer
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
		if d.GetType() == devices.COMSERVER {
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
		if d.GetType() == devices.LABDEVICE {
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
	/*switch device.GetType() {
	case devices.LABDEVICE:
		device = device.(devices.LabDevice)
	case devices.COMSERVER:
		device = device.(devices.ComServer)

	}*/
	return device
}
