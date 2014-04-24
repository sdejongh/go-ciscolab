package devices

import "net"

// ComServer Device
type ComServer struct {
	name       string
	ipAddress  net.IP
	telnetPort uint16
}

// Sets the name of the Comserver
func (c *ComServer) SetName(name string) {
	c.name = name
}

// Returns the name of the Comserver
func (c ComServer) GetName() string {
	return c.name
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
func NewComServer(name string, telnetPort uint16, ipAddress ...interface{}) *ComServer {
	if len(ipAddress) == 0 {
		return &ComServer{name: name, telnetPort: telnetPort, ipAddress: net.ParseIP("0.0.0.0")}
	}
	address := ipAddress[0]
	switch address.(type) {
	case string:
		return &ComServer{name: name, telnetPort: telnetPort, ipAddress: net.ParseIP(address.(string))}
	case net.IP:
		return &ComServer{name: name, telnetPort: telnetPort, ipAddress: address.(net.IP)}
	default:
		return &ComServer{name: "", telnetPort: 0, ipAddress: net.ParseIP("0.0.0.0")}
	}
}
