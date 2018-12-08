package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Guard struct {
	id      int
	time    int
	minutes map[int]int
}

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please enter an input file.")
		os.Exit(0)
	}

	inputFile := args[1]

	if inputFile == "" {
		fmt.Println("Please enter an input file.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}

	fileContents := string(bytes)

	array := strings.Split(fileContents, "\n")

	blankElementIndex := len(array) - 1

	listOfGuards := make(map[int]*Guard)

	evaluatedGuard := -1
	sleepBeginsTime := -1
	for i, str := range array {

		if i == blankElementIndex {
			continue
		}

		if strings.Contains(str, "begins shift") {

			matches := strings.Split(str, " ")
			idStr := strings.Trim(matches[3], "#")
			id, _ := strconv.Atoi(idStr)

			evaluatedGuard = id

		} else if strings.Contains(str, "falls asleep") {

			matches := strings.Split(str, " ")
			startTime := strings.Trim(matches[1], "]")
			hoursMinutes := strings.Split(startTime, ":")
			sleepStart, _ := strconv.Atoi(hoursMinutes[1])

			sleepBeginsTime = sleepStart

		} else if strings.Contains(str, "wakes up") {

			matches := strings.Split(str, " ")
			endTime := strings.Trim(matches[1], "]")
			hoursMinutes := strings.Split(endTime, ":")
			sleepEnd, _ := strconv.Atoi(hoursMinutes[1])

			// check if entry is nil
			if _, ok := listOfGuards[evaluatedGuard]; !ok {

				minutes := make(map[int]int)
				newGuard := &Guard{id: evaluatedGuard, time: 0, minutes: minutes}
				listOfGuards[evaluatedGuard] = newGuard
			}

			// increment the amount of time a guard is asleep
			listOfGuards[evaluatedGuard].time += (sleepEnd - sleepBeginsTime)

			for j := sleepBeginsTime; j < sleepEnd; j++ {
				listOfGuards[evaluatedGuard].minutes[j]++
			}
		}
	}

	biggestMinute := -1
	biggestMinuteAmount := -1
	maxId := -1

	for id, guard := range listOfGuards {
		for minute, val := range guard.minutes {
			if val > biggestMinuteAmount {
				biggestMinute = minute
				biggestMinuteAmount = val
				maxId = id
			}
		}
	}

	fmt.Println(maxId * biggestMinute)
}
