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

type State struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	Value    string
	Write    string
	Move     string
	Continue string
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

	beginStateRegex := regexp.MustCompile("Begin in state ([A-Z]).")

	performDiagnosisRegex := regexp.MustCompile(
		"Perform a diagnostic checksum after ([0-9]{1,16}) steps.")

	inStateRegex := regexp.MustCompile("In state ([A-Z]):")

	inIfTheCurrentRegex := regexp.MustCompile(
		"If the current value is ([0-9]):")

	writeTheValueRegex := regexp.MustCompile(
		"- Write the value ([0-9]).")

	moveOneSlotRegex := regexp.MustCompile(
		"- Move one slot to the ([a-z]{1,5}).")

	continueWithStateRegex := regexp.MustCompile(
		"- Continue with state ([A-Z]).")

	beginningState := ""
	diagAfter := 0
	parsingState := false
	turingStates := make(map[string]State)

	currentState := ""
	currentRule := Rule{}

	for i, l := range lines {

		line := strings.TrimSpace(l)

		if line == "" {
			parsingState = false
			currentState = ""
			continue

		} else if i == 0 {
			matches := beginStateRegex.FindAllStringSubmatch(line, -1)
			beginningState = matches[0][1]
			continue

		} else if i == 1 {
			matches := performDiagnosisRegex.FindAllStringSubmatch(line, -1)
			extractedInt, err := strconv.ParseInt(matches[0][1], 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			diagAfter = int(extractedInt)
			continue
		}

		//
		// Attempt to parse the state...
		//

		if parsingState {

			// If the current value is x.
			matches := inIfTheCurrentRegex.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 && len(matches[0]) > 1 {
				currentRule = Rule{matches[0][1], "", "", ""}
				continue
			}

			// Write the value x.
			matches = writeTheValueRegex.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 && len(matches[0]) > 1 {
				currentRule.Write = matches[0][1]
				continue
			}

			// Move one slot to the x.
			matches = moveOneSlotRegex.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 && len(matches[0]) > 1 {
				currentRule.Move = matches[0][1]
				continue
			}

			// Continue with state x.
			matches = continueWithStateRegex.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 && len(matches[0]) > 1 {
				currentRule.Continue = matches[0][1]
				tmpState := turingStates[currentState]
				tmpState.Rules = append(tmpState.Rules,
					currentRule)
				turingStates[currentState] = tmpState
				continue
			}

			continue
		}

		//
		// Start to parse the state...
		//

		matches := inStateRegex.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 && len(matches[0]) > 1 {
			currentState = matches[0][1]
			parsingState = true
			rules := make([]Rule, 0)
			turingStates[currentState] = State{
				currentState, rules}
		}
	}

	tape := make(map[int64]string)
	ongoingState := beginningState
	var pos int64 = 0

	for i := 0; i < diagAfter; i++ {

		thisState := turingStates[ongoingState]
		ruleForZero := Rule{}
		ruleForOne := Rule{}

		// golang arrays can be random, so determine which rule is
		// which before doing anything
		if thisState.Rules[0].Value == "0" {
			ruleForZero = thisState.Rules[0]
			ruleForOne = thisState.Rules[1]
		} else {
			ruleForZero = thisState.Rules[1]
			ruleForOne = thisState.Rules[0]
		}

		// handle the tape value...
		if tape[pos] == "" || tape[pos] == "0" {

			// if tape is empty or zero at this point
			tape[pos] = ruleForZero.Write
			if ruleForZero.Move == "left" {
				pos--
			} else {
				pos++
			}
			ongoingState = ruleForZero.Continue
			continue

		} else if tape[pos] == "1" {

			// if tape is one at this point
			tape[pos] = ruleForOne.Write
			if ruleForOne.Move == "left" {
				pos--
			} else {
				pos++
			}
			ongoingState = ruleForOne.Continue
			continue
		}
	}

	// checksum the tape
	var checksum int64 = 0
	for _, value := range tape {
		if value == "1" {
			checksum++
		}
	}

	fmt.Println("Part 1:", checksum)
}
