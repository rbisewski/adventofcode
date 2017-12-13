package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

	processes := make([][]int, 0)

	re := regexp.MustCompile("\\d+")

	// search thru each line
	for _, l := range lines {

		// obtain matches
		matches := re.FindAllString(l, -1)

		// convert matches
		subs := make([]int, 0)
		for _, m := range matches {
			num, err := strconv.ParseInt(m, 10, 64)
			if err != nil {
				os.Exit(1)
			}
			subs = append(subs, int(num))
		}

		// attach all except parent to the processes
		processes = append(processes, subs[1:])
	}

	count := make(map[int]bool)

	pipeToSubprocess(processes, count, 0)

	numOfPrograms := len(count)

	groups := 1

	// this could take awhile...
	for {
		// for every "process"
		for i := range processes {

			// skip if true
			if count[i] {
				continue
			}

			// else open a "pipe" to the "subprocess"
			pipeToSubprocess(processes, count, i)
			groups++
		}

		// if every single process has been explored at least once
		if len(count) >= len(processes) {
			break
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", numOfPrograms, groups)
}

func pipeToSubprocess(processes [][]int, count map[int]bool, i int) {

	count[i] = true

	for _, n := range processes[i] {

		// skip if true
		if count[n] {
			continue
		}

		pipeToSubprocess(processes, count, n)
	}
}
