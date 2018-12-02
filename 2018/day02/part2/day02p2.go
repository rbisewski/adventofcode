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

	for i, str1 := range array {

		if i == blankElementIndex {
			continue
		}

		runes1 := []rune(str1)

		for j, str2 := range array {

			if j == i {
				continue
			}
			if j == blankElementIndex {
				continue
			}

			lettersDifferent := 0

			runes2 := []rune(str2)

			for k := 0; k < len(runes1); k++ {

				if runes1[k] != runes2[k] {
					lettersDifferent++
				}

				if lettersDifferent > 1 {
					break
				}
			}

			if lettersDifferent == 1 {
				fmt.Println(str1, str2)
			}
		}
	}
}
