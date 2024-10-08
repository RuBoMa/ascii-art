package main

import (
	"asciiart/ascii"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a correct input go run . [INPUT]")
		return
	}
	input := os.Args[1]
	result := ascii.PrintAsciiArt(input)
	fmt.Print(result)
}
