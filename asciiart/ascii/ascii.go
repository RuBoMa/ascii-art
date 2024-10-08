package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(input string) string {
	inputfile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("error", err)
		return ""
	}

	result := ""

	if len(input) != 0 {
		//splitting the inputfile (standard.txt) into a slice by rows. Row 1 becomes index 0, row 2 index 1 etc.
		inputFileLines := strings.Split(string(inputfile), "\n")
		//splitting the input into a slice by new line
		words := strings.Split(input, "\\n")

		onlyNewLines := true

		//An empty index generates a newline, else the index is looped and matching indexes from inputFileLines slice (=rows) will be printed
		for _, word := range words {
			if word == "" {
				result += "\n"
			} else {
				onlyNewLines = false
				for i := 0; i < 8; i++ {
					for _, char := range word {
						result += inputFileLines[i+(int(char-32)*9)+1]
					}
					result += "\n"
				}
			}
		}
		//if the input consists only newlines, deducting the last empty index from the slice
		if onlyNewLines {
			result = result[:len(result)-1]
		}
	}

	return result
}
