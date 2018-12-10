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

	//
	// part2
	//

	polymerString = string(polymer)

	bestLength := len(polymerString)
	bestChar := 'a'

	for i := 'a'; i <= 'z'; i++ {

		lowercase := i
		uppercase := i - 32

		reduceString := strings.Replace(polymerString, string(lowercase), "", -1)
		reduceString = strings.Replace(reduceString, string(uppercase), "", -1)
		reduce := []rune(reduceString)

		for {
			unitsReplaced := false
			currentChar := '0'
			previousChar := '0'

			for i, char := range reduce {

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
					reduce = append(reduce[:i-1], reduce[i+1:]...)
					break
				}
			}

			if unitsReplaced == false {
				break
			}
		}

		if len(reduce) < bestLength {
			bestLength = len(reduce)
			bestChar = i
		}
	}

	fmt.Println(string(bestChar), bestLength)
}
