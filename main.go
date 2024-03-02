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
}
