package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

		keyParts := strings.Split(key, "/")

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

	shape := ".#./..#/###"

	for iteration := 1; iteration <= 5; iteration++ {

		if iteration == 1 {
			shape = arrangements[shape]
			continue
		}
	}

	hashtag := regexp.MustCompile("#")
	matches := hashtag.FindAllString(shape, -1)
	pixelsOn := len(matches)

	fmt.Println(shape, pixelsOn)
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

	lines := ""
	length := len(array[0])

	// for every character in the element-string...
	for i := 0; i < length; i++ {

		// obtain the i-th value for every element...
		line := ""
		for _, elm := range array {

			// ... append that character to the transposed matrix
			char := string(elm[i])
			line += char
		}
		lines += "/" + line
	}

	lines = strings.Trim(lines, "/")

	return lines
}
