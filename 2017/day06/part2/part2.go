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
	blockList = ""
)

func init() {
	flag.StringVar(&blockList, "file", "",
		"Enter the jump list filepath.")
}

func main() {

	flag.Parse()

	if blockList == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(blockList)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\t")

	if len(lines) < 1 {
		fmt.Println(0)
		os.Exit(0)
	}

	elements := make([]int, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		e, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			continue
		}
		elements = append(elements, int(e))
	}

	pos := 0
	steps := 0
	size := len(elements)
	cache := 0

	states := make(map[string]int, 0)

	count := 0
	for {

		if cache == 0 {

			newState := ""
			highest := -1
			for i, elm := range elements {
				if elm > highest {
					highest = elm
					pos = i
				}
				newState += "," + strconv.FormatInt(int64(elm), 10)
			}

			_, alreadyOccurred := states[newState]
			if alreadyOccurred {
				count = steps - states[newState]
				break
			}
			states[newState] = steps

			if highest == -1 {
				fmt.Println(0)
				os.Exit(1)
			}

			cache = elements[pos]
			elements[pos] = 0

			steps++

			continue
		}

		pos = (pos + 1) % size
		elements[pos]++
		cache--
	}

	fmt.Println(count)
}
