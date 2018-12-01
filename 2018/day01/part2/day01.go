package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	// splitting by \n tends to leave empty elements at the end
	blankElementIndex := len(array) - 1

	freq := 0

	freqMap := make(map[int]int)

	terminateNow := false

	numberOfLoops := 0

	for {
		for i, num := range array {

			if i == blankElementIndex {
				continue
			}

			number, _ := strconv.Atoi(num)
			freq += number

			freqMap[freq]++

			if freqMap[freq] > 1 {
				terminateNow = true
				break
			}
		}

		numberOfLoops++

		if terminateNow {
			break
		}
	}

	fmt.Println("The number of loops: ", numberOfLoops)
	fmt.Println("Twice frequency is: ", freq)
}
