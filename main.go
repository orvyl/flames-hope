package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const flamesHopeStr = "flameshope"

func setUp() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	setUp()

	fmt.Println("*** Welcome to Flames Hope! ***")

	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Print("Please enter your crush's name: ")
	var crushName string
	fmt.Scanf("%s", &crushName)

	result := "Walang forever:("

	randomResult := rand.Int()
	fmt.Println(randomResult)

	if randomResult % 2 == 0 {
		result = strconv.Itoa(rand.Intn(100)) + "%"
	}

	fmt.Println("As of", time.Now().Format(time.RFC1123), " you two are: " + result)
}
