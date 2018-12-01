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

	freq := 0

	for _, num := range array {
		number, _ := strconv.Atoi(num)
		freq += number
	}

	fmt.Println(freq)
}
