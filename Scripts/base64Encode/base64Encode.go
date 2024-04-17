package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Encode your code to base64: ")
	var code string
	fmt.Scan(&code)

	result := base64.StdEncoding.EncodeToString([]byte(code))

	fmt.Printf("Result: %v \n", result)
}
