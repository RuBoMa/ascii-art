package ascii

import (
	"fmt"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input string) string {

	var result string

	bannerFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("ERROR: Couldn't read the banner file: ", err)
	} else if len(input) != 0 {
		//splitting the banner file into a slice by rows (Index 0 = row 1)
		bannerFileLines := strings.Split(string(bannerFile), "\n")
		//splitting the input into a slice by literal new line
		words := strings.Split(input, "\\n")

		onlyNewLines := true

		//An empty index generates a newline, else the index is looped to match indexes from bannerFileLines
		for _, word := range words {
			if word == "" {
				result += "\n"
			} else {
				onlyNewLines = false
				for i := 1; i <= 8; i++ {
					for _, char := range word {
						result += bannerFileLines[i+(int(char-32)*9)]
					}
					result += "\n"
				}
			}
		}
		//if the input consists only newlines, deducting one newline
		if onlyNewLines && len(result) > 0 {
			result = result[1:]
		}
	}

	return result
}
