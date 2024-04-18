package main

import (
	"fmt"
	"myProj/myLibs"
)

func startProgram() {
	fmt.Println("Start program ...")
}

func main() {
	startProgram()
	myLibs.AddNumbers(2, 4)
	myLibs.DisplayMessage("Imported Message")
}
