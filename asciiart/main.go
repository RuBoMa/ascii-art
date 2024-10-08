package main

import (
	"asciiart/ascii"
	"fmt"
	"os"
)

func main() {

	if !isValid(os.Args) {
		os.Exit(2)
	}
	input := os.Args[1]
	result := ascii.PrintAsciiArt(input)
	fmt.Print(result)
}

func isValid(args []string) bool {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Invalid amount of arguments")
		return false
	}

	for _, r := range args[1] {
		if (r < 32 || r > 127) && r != '\n' {
			fmt.Printf("ERROR: %c is not a printable ascii character\n", r)
			return false
		}
	}
	return true
}
