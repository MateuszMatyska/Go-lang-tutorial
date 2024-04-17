package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Read data from file:")
	fmt.Println("")

	// open file
	f, err := os.Open("in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fmt.Printf("%v\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Save data to file:")
	fmt.Println("")

	f1, err1 := os.Create("out.txt")

	if err1 != nil {
		log.Fatal(err1)
	}

	defer f1.Close()

	buffer1 := bufio.NewWriter(f1)

	for _, line := range lines {
		_, err := buffer1.WriteString(line + "\n")
		if err != nil {
			log.Fatal((err))
		}
	}

	if err := buffer1.Flush(); err != nil {
		log.Fatal(err)
	}
}
