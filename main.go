package main

import "fmt"

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
	namesArray3 = append(namesArray3, "Michael")
	namesArray3 = append(namesArray3, "Jim")
	namesArray3 = append(namesArray3, "Dwight")
	namesArray3 = append(namesArray3, "Pam")
	fmt.Printf("Names 3: %v\n", namesArray3)
}
