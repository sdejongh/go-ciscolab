// CiscoLab project ciscolab.go
package main

import (
	"ciscolab/devices"
	"ciscolab/lab"
	"fmt"
)

func main() {

	c := devices.NewComServer("ComServer3", 23, "192.168.0.3")
	c2 := devices.NewComServer("ComServer1", 23, "192.168.0.1")
	c3 := devices.NewComServer("ComServer2", 23, "192.168.0.2")

	d := devices.NewLabDevice("P1R1", 33, c)
	d2 := devices.NewLabDevice("P1R3", 35, c)
	d3 := devices.NewLabDevice("P1R2", 34, c2)

	l := lab.NewLab()
	l.AddComServer(c, c2, c3)
	l.AddDevice(d, d2, d3)

	fmt.Println("CiscoLab GO version")

	fmt.Println(c.GetName(), "->", c.GetIpAddress(), ":", c.GetTelnetPort())
	fmt.Println(d.GetName(), "->", d.GetComServer().GetName(), ":", d.GetComServer().GetIpAddress(), ":", 2000+d.GetLine())

	fmt.Println(l.ListComServersByName(true))
	fmt.Println(l.ListDevicesByName(true))

	fmt.Println(l.IsComServerInLab("ComServer4"))
	fmt.Println(l.GetComServerByName("ComServer1"))

	fmt.Println(l.IsDeviceInLab("P1R1"))
	fmt.Println(l.GetDeviceByName("P1R1"))
}
