// CiscoLab project ciscolab.go
package main

import (
	//"ciscolab/devices"
	//"ciscolab/lab"
	"ciscolab/menu"
	"fmt"
)

func foo() {
	fmt.Println("Hello World!")
}

func main() {
	/*
		c := devices.NewComServer("ComServer3", 23, "192.168.0.3")
		c2 := devices.NewComServer("ComServer1", 23, "192.168.0.1")
		c3 := devices.NewComServer("ComServer2", 23, "192.168.0.2")

		d := devices.NewLabDevice("P1R1", 33, c)
		d2 := devices.NewLabDevice("P1R3", 35, c)
		d3 := devices.NewLabDevice("P1R2", 34, c2)

		l := lab.NewLab()
		l.AddDevice(d, d2, d3, c, c2, c3)
		l.RemoveDevice("P1R2")

		fmt.Println("CiscoLab GO version")

		fmt.Println(c.GetName(), "->", c.GetIpAddress(), ":", c.GetTelnetPort())
		fmt.Println(d.GetName(), "->", d.GetComServer().GetName(), ":", d.GetComServer().GetIpAddress(), ":", 2000+d.GetLine())

		fmt.Println("Is P1R1 in lab ?", l.IsDeviceInLab("P1R1"))

		fmt.Println(l.GetComServerNames(true))
		fmt.Println(l.GetLabDeviceNames(true))

		fmt.Println("P1R1 is at", l.GetDeviceByName("P1R1"))
	*/

	m := menu.NewMenu("CiscoLab v3.0alpha - GO version", menu.BORDER, 2, 1, 1, 2)
	m.SetWidth(80)
	m.AddTextLine("Simple text line left aligned", menu.TXTALIGN_LEFT)
	m.AddEmptyLine()
	m.AddTextLine("Simple text line centered", menu.TXTALIGN_CENTER)
	m.AddEmptyLine()
	m.AddTextLine("Simple text line right aligned", menu.TXTALIGN_RIGHT)
	m.AddEmptyLine()
	m.AddCommandLine(1, "Test command", foo, menu.TXTALIGN_LEFT)
	m.Display()
}
