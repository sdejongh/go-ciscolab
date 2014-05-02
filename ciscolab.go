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

	m := menu.NewMenu("CiscoLab v3.0alpha - GO version", "Your choice: ", menu.BORDER, 2, 1, 0, 2)
	m.SetWidth(50)
	m.AddTextLine("Left Aligned commands", menu.TXTALIGN_LEFT)
	m.AddEmptyLine()
	m.AddCommandLine(1, "Test command 1", foo, menu.TXTALIGN_LEFT)
	m.AddCommandLine(2, "Test command 2", foo, menu.TXTALIGN_LEFT)
	m.AddCommandLine(3, "Test command 3", foo, menu.TXTALIGN_LEFT)
	m.AddCommandLine(4, "Test command 4", foo, menu.TXTALIGN_LEFT)
	m.AddEmptyLine()
	m.AddTextLine("Right Aligned commands", menu.TXTALIGN_RIGHT)
	m.AddEmptyLine()
	m.AddCommandLine(5, "Test command 5", foo, menu.TXTALIGN_RIGHT)
	m.AddCommandLine(6, "Test command 6", foo, menu.TXTALIGN_RIGHT)
	m.AddCommandLine(7, "Test command 7", foo, menu.TXTALIGN_RIGHT)
	m.AddCommandLine(8, "Test command 8", foo, menu.TXTALIGN_RIGHT)
	m.Display()
}
