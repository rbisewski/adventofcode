package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	stream = ""
)

func init() {
	flag.StringVar(&stream, "file", "",
		"Enter the stream filepath.")
}

func main() {

	flag.Parse()

	if stream == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(stream)

	if err != nil {
		os.Exit(1)
	}

	contents := string(bytes)

	if contents == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	contents = strings.TrimSpace(contents)

	pos := 0
	end := len(contents)

	garbageCollecting := false

	strArray := make([]string, 0)
	for pos < end {
		currentChar := string(contents[pos])
		strArray = append(strArray, currentChar)
		pos++
	}

	pos = 0

	// clean up the garbage cancellation
	for pos < end {

		currentChar := strArray[pos]

		if currentChar == "!" && !garbageCollecting {
			garbageCollecting = true
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue

		} else if garbageCollecting {
			garbageCollecting = false
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue

		}

		pos++
	}

	garbageCollecting = false
	pos = 0

	// clean up the garbage between < and >
	for pos < end {

		currentChar := strArray[pos]

		if currentChar == "<" && !garbageCollecting {
			garbageCollecting = true
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue

		} else if currentChar != ">" && garbageCollecting {
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue

		} else if currentChar == ">" && garbageCollecting {
			garbageCollecting = false
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue
		}

		pos++
	}

	pos = 0

	// remove all non-bracket characters
	for pos < end {

		currentChar := strArray[pos]

		if currentChar != "{" && currentChar != "}" {
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue
		}

		pos++
	}

	pos = 0
	level := 0
	count := 0

	// add up the score
	for pos < end {

		currentChar := strArray[pos]

		if currentChar == "{" {
			level++
		} else if currentChar == "}" {
			count += level
			level--
		}

		pos++
	}

	// DEBUG: reassemble the string
	//output := ""
	//for _, str := range strArray {
	//	output += str
	//}
	//fmt.Println(output)

	fmt.Println(count)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
