package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const flamesHopeStr = "flameshope"

func main() {
	fmt.Println("*** Welcome to Flames Hope! ***")

	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Print("Please enter your crush's name: ")
	var crushName string
	fmt.Scanf("%s", &crushName)

	result := "Walang forever sa inyo :("

	if rand.Int31() % 2 == 0 {
		result = strconv.Itoa(rand.Intn(100)) + "%"
	}

	fmt.Println("RESULT: " + result)
}
