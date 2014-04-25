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
	l := lab.NewLab(3, 48)
	l.AddComServer(c, c2, c3)

	fmt.Println("CiscoLab GO version")

	fmt.Println(c.GetName(), "->", c.GetIpAddress(), ":", c.GetTelnetPort())
	fmt.Println(d.GetName(), "->", d.GetComServer().GetName(), ":", d.GetComServer().GetIpAddress(), ":", 2000+d.GetLine())

	fmt.Println(l.ListComServersByName(true))
}
