package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	polymerString := string(bytes)

	polymer := []rune(polymerString)

	for {
		unitsReplaced := 0
		currentChar := '0'
		previousChar := '0'

		for char := range currentPolymer {

			if currentChar == '0' {
				currentChar = char
				continue
			}

			previousChar = currentChar
			currentChar = char

			// TODO: check if the current char is same letter but opposite char, in which case trim the rune
		}

		if unitsReplaced == 0 {
			break
		}
	}

	fmt.Println(polymer)
}
