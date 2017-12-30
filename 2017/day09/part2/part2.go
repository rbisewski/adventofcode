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
	count := 0
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
			count++
			continue

		} else if currentChar == ">" && garbageCollecting {
			garbageCollecting = false
			strArray = remove(strArray, pos)
			end = len(strArray)
			continue
		}

		pos++
	}

	fmt.Println(count)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
