package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	// standard.txt (readData), split by '\n'

	// args, after main file
	args := os.Args[1]
	banner := os.Args[2]

	splitStr := strings.Split(string(readData(banner+".txt")), "\n")
	// args, split by new line
	splitArgs := strings.Split(string(args), "\\n")
	// for each slice of arg...
	for a := 0; a < len(splitArgs); a++ {

		runeArgs := []rune(splitArgs[a])
		// if the slice of arg contains nothing then print a new line
		if len(splitArgs[a]) == 0 {
			(print("\n"))
			// otherwise print the first line of each letter in the slice of arg, then 2nd - 8th
		} else {
			for i := 1; i <= 8; i++ {
				for j := 0; j <= len(runeArgs)-1; j++ {

					letterArgs := runeArgs[j]
					// rune to line number
					lineNumber := (int(letterArgs)-32)*9 + i
					// prints from line number to the next 8 lines
					fmt.Print(splitStr[lineNumber])
				}
				fmt.Println()
			}
		}
	}
}
