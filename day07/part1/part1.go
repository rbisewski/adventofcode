package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	towerList = ""
)

func init() {
	flag.StringVar(&towerList, "file", "",
		"Enter the tower list filepath.")
}

func main() {

	flag.Parse()

	if towerList == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(towerList)

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

	// obtain only the lines with arrows
	arrow_lines := make([]string, 0)
	for _, l := range lines {

		if strings.Contains(l, " -> ") {
			arrow_lines = append(arrow_lines, l)
		}
	}

	if len(arrow_lines) < 1 {
		fmt.Println("arrow_lines is empty")
		os.Exit(1)
	}

	bottomProgram := ""
	for _, l := range arrow_lines {

		elements := strings.Split(l, " ")

		if len(elements) < 1 {
			fmt.Println("elements is empty")
			os.Exit(1)
		}

		first := elements[0]

		programIsSub := false
		for _, l2 := range arrow_lines {

			if strings.Contains(l2, " "+first) {

				programIsSub = true
				break
			}
		}

		if !programIsSub {
			bottomProgram = first
			break
		}
	}

	fmt.Println(bottomProgram)
}
