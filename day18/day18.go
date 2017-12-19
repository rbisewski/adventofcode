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

	// ----------------------------------------------------------------

	zeroQueue := make([]int, 0)
	zeroRegisters := make(map[string]int)
	zeroRegisters["p"] = 0
	programZero = Program{0, zeroRegisters, zeroQueue, lines, 0, false,
		false}

	oneQueue := make([]int, 0)
	oneRegisters := make(map[string]int)
	oneRegisters["p"] = 1
	programOne = Program{1, oneRegisters, oneQueue, lines, 0, false,
		false}

	for {
		programZero = executeNext(programZero)
		programOne = executeNext(programOne)

		if programZero.Waiting && programOne.Waiting {
			break
		}
	}

	fmt.Println("Part 2:", programOneSentValue)
}

func executeNext(p Program) Program {

	if p.Done {
		return p

	} else if p.Pos < 0 || p.Pos >= len(p.Lines) {
		p.Waiting = true
		p.Done = true
		return p
	}

	re := regexp.MustCompile("[a-z]")

	instruct := strings.Split(p.Lines[p.Pos], " ")

	first := 0
	second := 0
	firstStr := ""

	firstStr = re.FindString(instruct[1])
	if firstStr == "" {
		x, _ := strconv.ParseInt(instruct[1], 10, 64)
		first = int(x)
	} else {
		first = p.Registers[instruct[1]]
	}

	if len(instruct) > 2 {

		if re.FindString(instruct[2]) == "" {
			x, _ := strconv.ParseInt(instruct[2], 10, 64)
			second = int(x)
		} else {
			second = p.Registers[instruct[2]]
		}
	}

	if p.Waiting && instruct[0] != "rcv" {
		return p
	}

	switch instruct[0] {

	case "snd":
		if p.ID == 0 {
			programOne.Queue = append(programOne.Queue, first)
		} else if p.ID == 1 {
			programOneSentValue++
			programZero.Queue = append(programZero.Queue, first)
		}
	case "set":
		p.Registers[firstStr] = second
	case "add":
		p.Registers[firstStr] += second
	case "mul":
		p.Registers[firstStr] *= second
	case "mod":
		p.Registers[firstStr] = p.Registers[firstStr] % second
	case "rcv":
		if len(p.Queue) > 0 {
			value := 0
			value, p.Queue = p.Queue[0], p.Queue[1:]
			p.Registers[firstStr] = value
			p.Waiting = false
		} else {
			p.Waiting = true
			return p
		}
	case "jgz":
		if first > 0 {
			p.Pos = p.Pos + second
			return p
		}
	}

	p.Pos = p.Pos + 1

	return p
}
