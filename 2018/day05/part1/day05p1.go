package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please enter an input file.")
		os.Exit(0)
	}

	inputFile := args[1]

	if inputFile == "" {
		fmt.Println("Please enter an input file.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}

	polymerString := strings.TrimSpace(string(bytes))

	polymer := []rune(polymerString)

	for {
		unitsReplaced := false
		currentChar := '0'
		previousChar := '0'

		for i, char := range polymer {

			if currentChar == '0' {
				currentChar = char
				continue
			}

			previousChar = currentChar
			currentChar = char

			// check if the current char is same letter but opposite char, in which case trim the rune
			if previousChar+32 == currentChar {
				unitsReplaced = true

			} else if previousChar-32 == currentChar {
				unitsReplaced = true
			}

			if unitsReplaced {
				polymer[i] = '0'
				polymer[i-1] = '0'
				break
			}
		}

		polymerString = string(polymer)
		polymerString = strings.Replace(polymerString, "00", "", -1)
		polymer = []rune(polymerString)

		if unitsReplaced == false {
			break
		}
	}

	fmt.Println(len(polymer))
}
