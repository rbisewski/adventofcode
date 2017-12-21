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

	if len(lines) == 0 {
		os.Exit(1)
	}

	// read in input
	arrangements := make(map[string]string)
	for _, line := range lines {

		elements := strings.Split(line, " => ")

		if len(elements) != 2 {
			continue
		}

		key := strings.TrimSpace(elements[0])
		value := strings.TrimSpace(elements[1])

		arrangements[key] = value
	}

	// for every key, flip each of the input keys to account for all
	// possibilities
	for key, value := range arrangements {

		keyParts := strings.Split(key, "\\")

		// vertical flipped
		verticalFlippedKey := ""
		for _, elm := range keyParts {
			verticalFlippedKey += verticalFlipString(elm) + "/"
		}
		verticalFlippedKey = strings.Trim(verticalFlippedKey, "/")
		arrangements[verticalFlippedKey] = value

		// 90 degrees flipped
		nintyDegreesFlippedKey := nintyDegreeFlipString(keyParts)
		arrangements[nintyDegreesFlippedKey] = value

		// horizontal flipped
		horizontalFlippedKey := ""
		for _, elm := range keyParts {
			horizontalFlippedKey = elm + "/" + horizontalFlippedKey
		}
		horizontalFlippedKey = strings.Trim(horizontalFlippedKey, "/")
		arrangements[horizontalFlippedKey] = value
	}

	fmt.Println(arrangements)
}

func verticalFlipString(s string) string {

	if len(s) < 2 {
		return s
	}

	newString := ""

	for i := len(s) - 1; i >= 0; i-- {
		char := string(s[i])
		newString += char
	}

	return newString
}

func nintyDegreeFlipString(elements []string) string {

	if len(elements) < 1 {
		return ""
	}

	// inverse the position the elements
	array := make([]string, 0)
	for i := len(elements) - 1; i >= 0; i-- {
		array = append(array, elements[i])
	}

	// prepare each of the subsections
	matrix := make([][]string, 0)
	for i := 0; i < len(array); i++ {

		subsection := make([]string, 0)

		matrix = append(matrix, subsection)
	}

	// for every element...
	// TODO: fix this
	for _, elm := range array {

		// ... for every character in the element-string
		for i := 0; i < len(elm); i++ {

			// ... append that character to the transposed matrix
			char := string(elm[i])
			matrix[i] = append(matrix[i], char)
		}
	}

	newString := ""
	for _, tuple := range matrix {

		line := ""
		for _, elm := range tuple {
			line += elm
		}
		newString += line + "/"
	}
	newString = strings.Trim(newString, "/")

	return newString
}
