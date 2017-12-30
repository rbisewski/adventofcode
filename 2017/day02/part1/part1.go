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
	spreadsheet = ""
)

func init() {
	flag.StringVar(&spreadsheet, "file", "",
		"Enter the spreadsheet filepath.")
}

func main() {

	flag.Parse()

	if spreadsheet == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(spreadsheet)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\n")

	var value int64 = 0

	for _, line := range lines {

		elements := strings.Split(line, "\t")

		elementsAsInts := make([]int64, 0)
		for _, e := range elements {

			if e == "" {
				continue
			}

			res, err := strconv.ParseInt(e, 10, 64)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			elementsAsInts = append(elementsAsInts, res)
		}

		if len(elementsAsInts) < 2 {
			continue
		}

		largest := elementsAsInts[0]
		smallest := elementsAsInts[0]

		for _, e := range elementsAsInts {
			if e > largest {
				largest = e
			}
			if e < smallest {
				smallest = e
			}
		}

		value += (largest - smallest)
	}

	fmt.Println(value)
}
