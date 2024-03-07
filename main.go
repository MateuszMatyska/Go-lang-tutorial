package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferanceName string = "Go lang Conf. 2024"
	const totalTickets int = 50
	var remainingTickets int = totalTickets

	fmt.Printf("Welcome on %v \n", conferanceName)
	fmt.Printf("We have total of %v tickets and, %v are still available \n", totalTickets, remainingTickets)
	fmt.Println("Get your ticket here")

	var userName string
	fmt.Scan(&userName)

	banner := "Welcome"
	fmt.Printf("%v, %v to our %v \n", userName, banner, conferanceName)

	var namesArray1 = [3]string{"Tom", "Arthur", "John"}
	fmt.Printf("Names 1: %v\n", namesArray1)

	var namesArray2 [3]string
	namesArray2[0] = "Elliot"
	namesArray2[1] = "Angela"
	namesArray2[2] = "Darlene"
	fmt.Printf("Names 2: %v\n", namesArray2)

	var namesArray3 []string
	namesArray3 = append(namesArray3, "Michael Scott")
	namesArray3 = append(namesArray3, "Jim Halpert")
	namesArray3 = append(namesArray3, "Dwight Schrute")
	namesArray3 = append(namesArray3, "Pam Beesly")
	fmt.Printf("Names 3: %v\n", namesArray3)

	for index, person := range namesArray3 {
		name := strings.Fields(person)[0]
		number := index + 1
		fmt.Printf("[%v]: Person Name: %v \n", number, name)
	}

	for _, person := range namesArray3 {
		name := strings.Fields(person)[1]
		fmt.Printf("Person Last Name: %v \n", name)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	const n = "GoLang"

	for index, l := range n {
		letter := string(l)
		fmt.Printf("[%v]{%v}: %v\n", index, l, letter)
	}
}
