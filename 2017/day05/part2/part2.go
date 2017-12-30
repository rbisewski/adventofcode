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
	jump_list = ""
)

func init() {
	flag.StringVar(&jump_list, "file", "",
		"Enter the jump list filepath.")
}

func main() {

	flag.Parse()

	if jump_list == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(jump_list)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\n")

	var steps int64 = 0

	elements := make([]int64, 0)
	for _, line := range lines {
		e, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			continue
		}
		elements = append(elements, e)
	}

	var pos int64 = 0
	jump := elements[pos]
	for {
		if pos < 0 || pos >= int64(len(elements)) {
			break
		}
		jump = elements[pos]
		if elements[pos] > 2 {
			elements[pos]--
		} else {
			elements[pos]++
		}
		pos += jump
		steps++
	}

	fmt.Println(steps)
}
