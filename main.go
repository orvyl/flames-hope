package main

import (
	"fmt"
	"github.com/orvyl/flames-hope/util"
)

func main() {
	fmt.Println("*** Welcome to Flames Hope! ***")

	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Print("Please enter your crush's name: ")
	var crushName string
	fmt.Scanf("%s", &crushName)

	fmt.Println("RESULT: " + util.Compute(name, crushName))
}
