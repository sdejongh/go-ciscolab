// CiscoLab project ciscolab.go
package main

import (
	"ciscolab/lab"
	//"ciscolab/network"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

func foo() {
	fmt.Println("Hello World!")
}

type User struct {
	Username string
	Password []byte
}

func main() {
	lab := lab.NewLab()
	lab.BuildFromPackage()

	file, err := os.Create("users.dat")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := gob.NewEncoder(file)

	hasher := md5.New()

	io.WriteString(hasher, "testpwd1")

	u1 := User{Username: "test1", Password: hasher.Sum(nil)}

	io.WriteString(hasher, "testpwd2")

	u2 := User{Username: "test2", Password: hasher.Sum(nil)}

	users := make([]User, 0)
	users = append(users, u1, u2)

	encoder.Encode(users)

	var u []User

	infile, err := os.Open("users.dat")
	decoder := gob.NewDecoder(infile)
	err2 := decoder.Decode(&u)
	if err2 != nil {
		fmt.Println("Error decoding:", err)
	}
	fmt.Println(u)

	fmt.Printf("%d", hasher.Sum(nil))

	//sshServer := network.NewSSHServer("0.0.0.0:9022", "id_rsa")
	//sshServer.Listen()

}
