package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputFile = ""
)

func init() {
	flag.StringVar(&inputFile, "inputFile", "",
		"Enter the input file location.")
}

func main() {

	flag.Parse()

	if inputFile == "" {
		fmt.Println("No file was specified.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}

	fileContents := string(bytes)

	array := strings.Split(fileContents, "\n")

	blankElementIndex := len(array) - 1

	numberTwice := 0
	numberTriple := 0

	for i, str := range array {

		if i == blankElementIndex {
			continue
		}

		hasTwiceBeenIncremented := false
		hasTripleBeenIncremented := false

		runes := []rune(str)
		for _, r := range runes {
			char := string(r)
			count := strings.Count(str, char)

			if count == 2 {
				hasTwiceBeenIncremented = true
			}
			if count == 3 {
				hasTripleBeenIncremented = true
			}
		}

		if hasTwiceBeenIncremented {
			numberTwice++
		}
		if hasTripleBeenIncremented {
			numberTriple++
		}
	}

	fmt.Println(numberTwice * numberTriple)
}
