package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Decode base64 code:")
	var base64Code string
	fmt.Scan(&base64Code)

	result, error := base64.StdEncoding.DecodeString(base64Code)

	if error != nil {
		fmt.Println("Something went wrong")
	} else {
		fmt.Printf("Result: %v\n", string(result))
	}

}
