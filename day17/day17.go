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

type Command struct {
	Action string
	First  int64
	Second int64
}

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

	stepsInt64, err := strconv.ParseInt(contents, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	steps := int(stepsInt64)

	array := make([]int, 0)
	pos := 0

	array = append(array, 0)
	array = append(array, 1)

	max := 2017

	for i := 2; i <= max; i++ {

		pos += steps
		pos %= len(array)

		// insert
		array = append(array, 0)
		copy(array[pos+1:], array[pos:])
		array[pos] = i
		pos++
	}

	valueAfter2017 := 0
	for i, elm := range array {

		if elm == 2017 {
			valueAfter2017 = array[i+1]
		}
	}

	fmt.Println("Part 1:", valueAfter2017)

	valueAfter50000000 := 0

	// treat the puzzle as a big combination lock that cycles around
	ptr := 0
	for i := 1; i <= 50000000; i++ {
		ptr = (ptr+steps)%i + 1
		if ptr == 1 {
			valueAfter50000000 = i
		}
	}

	fmt.Println("Part 2:", valueAfter50000000)
}
