package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

	dirtyBytes, err := ioutil.ReadFile(file)
	if err != nil || len(dirtyBytes) < 1 {
		os.Exit(1)
	}

	filestr := string(dirtyBytes)

	if filestr == "" {
		os.Exit(1)
	}

	filestr = strings.TrimSpace(filestr)

	if filestr == "" {
		os.Exit(1)
	}

	bytes := []byte(filestr)

	addendum := []byte{17, 31, 73, 47, 23}

	for _, a := range addendum {
		bytes = append(bytes, a)
	}

	// generate list of numbers from 0 to 255
	list := make([]int64, 0)
	var counter int64 = 0
	for counter < 256 {
		list = append(list, counter)
		counter++
	}

	// convert bytes to int lengths
	lengths := make([]int64, 0)
	for _, b := range bytes {
		newLength := int64(b)
		lengths = append(lengths, newLength)
	}

	var pos int64 = 0
	var skip int64 = 0

	for i := 0; i < 64; i++ {
		list, pos, skip = executeRound(list, lengths, pos, skip)
	}

	hashnik, err := densifyHash(list)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	finalHash := ""
	for _, decimal := range hashnik {
		hex := fmt.Sprintf("%x", decimal)
		if len(hex) == 1 {
			hex = "0" + hex
		}
		finalHash += hex
	}

	fmt.Println(finalHash)
	os.Exit(0)
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

func executeRound(list []int64, lengths []int64, pos int64, skip int64) ([]int64,
	int64, int64) {

	var listSize int64 = int64(len(list))

	for _, l := range lengths {

		list = flipElements(list, pos, l)
		pos = (pos + l + skip) % listSize
		skip++
	}

	return list, pos, skip
}

func densifyHash(list []int64) ([]int64, error) {

	if len(list) < 1 {
		return nil, fmt.Errorf("invalid or empty list")
	}

	remainder := len(list) % 16

	if remainder != 0 {
		return nil, fmt.Errorf("list is not a multiple of 16")
	}

	newList := make([]int64, 0)

	var element int64 = 0
	count := 1
	for _, l := range list {

		if count == 1 {
			element = l
			count++
			continue
		}

		element ^= l

		if count == 16 {
			count = 1
			newList = append(newList, element)
			continue
		}

		count++
	}

	return newList, nil
}
