package main

import (
	"fmt"
	"mod_pack/greeting"
	"mod_pack/user"

	"github.com/k0kubun/pp"
)

func main() {
	greeting.SayHello()
	greeting.SayFuckYuo()

	smt := user.NewUser("Dastan", 23)

	fmt.Println(smt.GetName())
	smt.SetNewName("Vika")

	pp.Println(smt)

}
