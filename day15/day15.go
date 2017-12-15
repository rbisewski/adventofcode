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
	file = ""
)

func init() {
	flag.StringVar(&file, "file", "",
		"Enter the file filepath.")
}

func main() {

	flag.Parse()

	if file == "" {
		fmt.Println("Please enter a valid file.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		os.Exit(1)
	}

	contents := string(bytes)

	if contents == "" {
		os.Exit(1)
	}

	contents = strings.TrimSpace(contents)

	if contents == "" {
		os.Exit(1)
	}

	lines := strings.Split(contents, "\n")

	if len(lines) < 1 {
		os.Exit(1)
	}

	var generatorAFactor int64 = 16807
	var generatorAPrevious int64 = 16807

	var generatorBFactor int64 = 48271
	var generatorBPrevious int64 = 48271

	var generationDividend int64 = 2147483647

	for i, l := range lines {

		elements := strings.Split(l, " ")

		if len(elements) < 5 || elements[4] == "" {
			os.Exit(1)
		}

		castedInt, err := strconv.ParseInt(elements[4], 10, 64)
		if err != nil {
			os.Exit(1)
		}

		if i == 0 {
			generatorAPrevious = castedInt
		} else if i == 1 {
			generatorBPrevious = castedInt
		}
	}

	count := 0
	judgmentAmount := 40000000

	for i := 0; i < judgmentAmount; i++ {

		generatorA := (generatorAPrevious * generatorAFactor) % generationDividend
		generatorB := (generatorBPrevious * generatorBFactor) % generationDividend

		generatorAPrevious = generatorA
		generatorBPrevious = generatorB

		gaAsString := strconv.FormatInt(generatorAPrevious, 2)
		gbAsString := strconv.FormatInt(generatorBPrevious, 2)

		ga16 := last16Chars(gaAsString)
		gb16 := last16Chars(gbAsString)

		if ga16 == gb16 {
			count++
		}
	}

	fmt.Println(count)
}

func last16Chars(str string) string {

	if str == "" {
		return ""
	}

	end := len(str) - 1
	start := end - 15

	if start < 0 {
		start = 0
	}

	last16 := ""
	for i := start; i <= end; i++ {

		char := string(str[i])
		last16 += char
	}

	return last16
}
