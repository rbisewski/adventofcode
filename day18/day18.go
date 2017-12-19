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
)

type Command struct {
	Action string
	First  int64
	Second int64
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

	recoverIsFirstNonZero := false

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

		case "snd":
			lastSound = first
		case "set":
			registers[firstStr] = second
		case "add":
			registers[firstStr] += second
		case "mul":
			registers[firstStr] *= second
		case "mod":
			registers[firstStr] = registers[firstStr] % second
		case "rcv":
			if registers[firstStr] != 0 {
				recoverIsFirstNonZero = true
			}
		case "jgz":
			if first > 0 {
				i += second
				continue
			}
		}

		if recoverIsFirstNonZero {
			break
		}

		i++
	}

	if recoverIsFirstNonZero {
		fmt.Println("Part 1:", lastSound)
	} else {
		fmt.Println("Part 1:", "rcv is never zero")
	}
}
