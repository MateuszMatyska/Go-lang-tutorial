package main

import (
	"fmt"
)

func main() {
	fmt.Println("Input message: ")
	var message string
	var encryptedMessage string
	fmt.Scan(&message)

	for _, letter := range message {
		number := int(letter)
		encryptedMessage = encryptedMessage + string(rune(number+3))
	}

	fmt.Printf("Encrypted Message: %v\n", encryptedMessage)
}
