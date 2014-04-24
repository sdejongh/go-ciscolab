// CiscoLab project ciscolab.go
package main

import (
	"ciscolab/devices"
	"fmt"
)

func main() {

	c := devices.NewComServer("ComServer1", 23, "192.168.0.1")
	d := devices.NewLabDevice("P1R1", 33, c)

	fmt.Println("CiscoLab GO version")

	fmt.Println(c.GetName(), "->", c.GetIpAddress(), ":", c.GetTelnetPort())
	fmt.Println(d.GetName(), "->", d.GetComServer().GetName(), ":", d.GetComServer().GetIpAddress(), ":", 2000+d.GetLine())

}
