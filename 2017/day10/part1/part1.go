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
		fmt.Println(0)
		os.Exit(0)
	}

	contents = strings.TrimSpace(contents)

	list := make([]int64, 0)
	var counter int64 = 0
	for counter < 256 {
		list = append(list, counter)
		counter++
	}

	lengthsAsString := strings.Split(contents, ",")

	lengths := make([]int64, 0)
	for _, length := range lengthsAsString {
		l, err := strconv.ParseInt(length, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		lengths = append(lengths, l)
	}

	var listSize int64 = int64(len(list))
	var result int64 = 0
	var pos int64 = 0
	var skip int64 = 0

	for _, l := range lengths {

		list = flipElements(list, pos, l)
		pos = (pos + l + skip) % listSize
		skip++
	}

	result = list[0] * list[1]
	fmt.Println(result)
}

func flipElements(list []int64, pos int64, length int64) []int64 {

	if length < 2 {
		return list
	}

	start := pos
	end := (pos + length) % int64(len(list))
	pos2 := start

	// obtain slice
	slice := make([]int64, 0)
	for i := start; i != end; i = (i + 1) % int64(len(list)) {
		slice = append(slice, list[i])
	}

	// inverse swap
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// re-insert
	for _, element := range slice {
		list[pos2] = element
		pos2 = (pos2 + 1) % int64(len(list))
	}

	return list
}
