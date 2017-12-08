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
	jumpInstructs = ""
)

type Instruction struct {
	Register         string
	Increment        bool
	Value            int64
	FirstComparator  string
	Operator         string
	SecondComparator int64
}

func init() {
	flag.StringVar(&jumpInstructs, "file", "",
		"Enter the tower list filepath.")
}

func main() {

	flag.Parse()

	if jumpInstructs == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(jumpInstructs)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\n")

	if len(lines) < 1 {
		fmt.Println("lines are empty")
		os.Exit(0)
	}

	listOfInstructions := make([]Instruction, 0)
	for _, l := range lines {

		if l == "" {
			continue
		}

		elements := strings.Split(l, " ")

		if len(elements) != 7 {
			fmt.Println("invalid instruction: ", l)
			os.Exit(1)
		}

		register := elements[0]

		increment := true
		if elements[1] == "inc" {
			increment = true
		} else if elements[1] == "dec" {
			increment = false
		}

		value, err := strconv.ParseInt(elements[2], 10, 64)
		if err != nil {
			fmt.Println("invalid integer: ", elements[2])
			os.Exit(1)
		}

		firstComparator := elements[4]
		operator := elements[5]

		secondComparator, err := strconv.ParseInt(elements[6], 10,
			64)
		if err != nil {
			fmt.Println("invalid integer: ", elements[6])
			os.Exit(1)
		}

		newInstr := Instruction{register, increment, value,
			firstComparator, operator, secondComparator}

		listOfInstructions = append(listOfInstructions, newInstr)
	}

	// map of registers
	regi := make(map[string]int64)

	for _, instr := range listOfInstructions {

		doOperation := false

		_, alreadyExists := regi[instr.Register]
		if !alreadyExists {
			regi[instr.Register] = 0
		}

		_, alreadyExists = regi[instr.FirstComparator]
		if !alreadyExists {
			regi[instr.FirstComparator] = 0
		}

		switch instr.Operator {
		case ">":
			if regi[instr.FirstComparator] >
				instr.SecondComparator {
				doOperation = true
			}
		case ">=":
			if regi[instr.FirstComparator] >=
				instr.SecondComparator {
				doOperation = true
			}
		case "<":
			if regi[instr.FirstComparator] <
				instr.SecondComparator {
				doOperation = true
			}
		case "<=":
			if regi[instr.FirstComparator] <=
				instr.SecondComparator {
				doOperation = true
			}
		case "==":
			if regi[instr.FirstComparator] ==
				instr.SecondComparator {
				doOperation = true
			}
		case "!=":
			if regi[instr.FirstComparator] !=
				instr.SecondComparator {
				doOperation = true
			}
		}

		// move on to the next instruction
		if !doOperation {
			continue
		}

		if instr.Increment {
			regi[instr.Register] += instr.Value
		} else {
			regi[instr.Register] -= instr.Value
		}
	}

	// TODO: improve this
	var highest int64 = -1000000000000
	for _, r := range regi {
		if r > highest {
			highest = r
		}
	}

	fmt.Println(highest)
}
