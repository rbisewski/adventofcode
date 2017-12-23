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

	lastSound = 0

	programOneSentValue = 0

	programZero Program
	programOne  Program
)

type Program struct {
	ID        int
	Registers map[string]int
	Queue     []int
	Lines     []string
	Pos       int
	Waiting   bool
	Done      bool
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

	lines := strings.Split(contents, "\n")

	if len(lines) == 0 {
		os.Exit(1)
	}

	//
	// Part 1
	//

	mulInstructionsExecuted := 0

	registers := make(map[string]int)

	re := regexp.MustCompile("[a-z]")

	i := 0

	for i >= 0 && i < len(lines) {

		instruct := strings.Split(lines[i], " ")

		first := 0
		second := 0
		firstStr := ""

		firstStr = re.FindString(instruct[1])
		if firstStr == "" {
			x, _ := strconv.ParseInt(instruct[1], 10, 64)
			first = int(x)
		} else {
			first = registers[instruct[1]]
		}

		if len(instruct) > 2 {

			if re.FindString(instruct[2]) == "" {
				x, _ := strconv.ParseInt(instruct[2], 10, 64)
				second = int(x)
			} else {
				second = registers[instruct[2]]
			}
		}

		switch instruct[0] {

		case "set":
			registers[firstStr] = second
		case "add":
			registers[firstStr] += second
		case "sub":
			registers[firstStr] -= second
		case "mul":
			registers[firstStr] *= second
			mulInstructionsExecuted++
		case "jnz":
			if first != 0 {
				i += second
				continue
			}
		}

		i++
	}

	fmt.Println("Part 1:", mulInstructionsExecuted)

	//
	// Part 2
	//

	registers = make(map[string]int)

	registers["a"]++

	i = 0

	for i >= 0 && i < len(lines) {

		instruct := strings.Split(lines[i], " ")

		first := 0
		second := 0
		firstStr := ""

		firstStr = re.FindString(instruct[1])
		if firstStr == "" {
			x, _ := strconv.ParseInt(instruct[1], 10, 64)
			first = int(x)
		} else {
			first = registers[instruct[1]]
		}

		if len(instruct) > 2 {

			if re.FindString(instruct[2]) == "" {
				x, _ := strconv.ParseInt(instruct[2], 10, 64)
				second = int(x)
			} else {
				second = registers[instruct[2]]
			}
		}

		switch instruct[0] {

		case "set":
			registers[firstStr] = second
		case "add":
			registers[firstStr] += second
		case "sub":
			registers[firstStr] -= second
		case "mul":
			registers[firstStr] *= second
			mulInstructionsExecuted++
		case "jnz":
			if first != 0 {
				i += second
				continue
			}
		}

		i++
	}

	fmt.Println("Part 2:", registers["h"])
}
