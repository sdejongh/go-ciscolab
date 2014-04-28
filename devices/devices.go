package devices

import (
	"net"
)

const (
	LABDEVICE DeviceType = iota
	COMSERVER DeviceType = iota
)

/* TYPE DEFINITIONS */

// Type of device, ComServer or LabDevice
type DeviceType int

// Generic Device Interface
type Devicer interface {
	GetName() string
	SetName(name string)
	GetType() DeviceType
}

// ComServer Device
type ComServer struct {
	name       string
	ipAddress  net.IP
	telnetPort uint16
}

// Lab Device
type LabDevice struct {
	name      string
	comserver *ComServer
	line      int16
}

/* COMSERVER related functions */

// Sets the name of the Comserver
func (c *ComServer) SetName(name string) {
	c.name = name
}

// Returns the name of the Comserver
func (c ComServer) GetName() string {
	return c.name
}

// Returns ComServer Type
func (d ComServer) GetType() DeviceType {
	return COMSERVER
}

// Sets the IP address of the Comserver
func (c *ComServer) SetIpAddress(ipAddress net.IP) {
	c.ipAddress = ipAddress
}

// Returns the IP address of the Comserver
func (c ComServer) GetIpAddress() net.IP {
	return c.ipAddress
}

// Sets the TCP port used for telnet connection
func (c *ComServer) SetTelnetPort(port uint16) {
	c.telnetPort = port
}

// Returns the TCP port used for telnet connections
func (c ComServer) GetTelnetPort() uint16 {
	return c.telnetPort
}

// ComServer Type Constructor
func NewComServer(name string, telnetPort uint16, ipAddress interface{}) *ComServer {
	switch ipAddress.(type) {
	case string:
		return &ComServer{name: name, telnetPort: telnetPort, ipAddress: net.ParseIP(ipAddress.(string))}
	case net.IP:
		return &ComServer{name: name, telnetPort: telnetPort, ipAddress: ipAddress.(net.IP)}
	default:
		return &ComServer{name: "", telnetPort: 0, ipAddress: net.ParseIP("0.0.0.0")}
	}
}

/* LABDEVICE related functions */

// Sets the name of the LabDevice
func (d *LabDevice) SetName(name string) {
	d.name = name
}

// Returns the name of the LavDevice
func (d LabDevice) GetName() string {
	return d.name
}

// Returns LabDevice Type
func (d LabDevice) GetType() DeviceType {
	return LABDEVICE
}

// Sets the comserver to which the LabDevice is connected through console port
func (d *LabDevice) SetComServer(comserver *ComServer) {
	d.comserver = comserver
}

// Returns the comserver to which the LabDevice is connected through the console port
func (d LabDevice) GetComServer() *ComServer {
	return d.comserver
}

// Sets the terminal line number of the comserver connected to the LabDevice
func (d *LabDevice) SetLine(line int16) {
	d.line = line
}

// Returns the line number of the comserver connected to the LabDevice
func (d LabDevice) GetLine() int16 {
	return d.line
}

// LabDevice type constructor
func NewLabDevice(name string, line int16, comserver *ComServer) *LabDevice {
	return &LabDevice{name: name, line: line, comserver: comserver}
}
