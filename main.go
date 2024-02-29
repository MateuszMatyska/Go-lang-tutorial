package main

import "fmt"

func main() {
	const conferanceName = "Go lang Conf. 2024"
	const totalTickets = 50
	var remainingTickets = totalTickets

	fmt.Printf("Welcome on %v \n", conferanceName)
	fmt.Printf("We have total of %v tickets and, %v are still available \n", totalTickets, remainingTickets)
	fmt.Println("Get your ticket here")

}
