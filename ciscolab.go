// CiscoLab project ciscolab.go
package main

import (
	//"ciscolab/devices"
	"ciscolab/lab"
	//"ciscolab/menu"
	"fmt"
)

func foo() {
	fmt.Println("Hello World!")
}

func main() {
	lab := lab.NewLab()
	lab.BuildFromPackage()
	fmt.Println(lab.GetComServerNames(true))
	fmt.Println(lab.GetLabDeviceNames(true))

}
