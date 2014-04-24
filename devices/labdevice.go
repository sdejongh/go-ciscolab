package devices

// Lab Device
type LabDevice struct {
	name      string
	comserver *ComServer
	line      int16
}

// Sets the name of the device
func (d *LabDevice) SetName(name string) {
	d.name = name
}

// Returns the name of the device
func (d LabDevice) GetName() string {
	return d.name
}

// Sets the comserver to which the divice is connected through console port
func (d *LabDevice) SetComServer(comserver *ComServer) {
	d.comserver = comserver
}

// Returns the comserver to which the device is connected through the console port
func (d LabDevice) GetComServer() *ComServer {
	return d.comserver
}

// Sets the terminal line number of the comserver connected to the device
func (d *LabDevice) SetLine(line int16) {
	d.line = line
}

// Returns the line number of the comserver connected to the device
func (d LabDevice) GetLine() int16 {
	return d.line
}

// LabDevice type constructor
func NewLabDevice(name string, line int16, comserver *ComServer) *LabDevice {
	return &LabDevice{name: name, line: line, comserver: comserver}
}
