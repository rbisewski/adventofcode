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

	cmdStrings := strings.Split(contents, ",")

	if len(cmdStrings) < 1 {
		os.Exit(1)
	}

	spin := regexp.MustCompile("s(\\d{1,2})")
	exchange := regexp.MustCompile("x(\\d{1,2})/(\\d{1,2})")
	partner := regexp.MustCompile("p([a-p])/([a-p])")

	cmds := make([]Command, 0)

	//
	// Construct all of the commands
	//

	for _, cmdStr := range cmdStrings {

		matches := spin.FindStringSubmatch(cmdStr)
		if len(matches) == 2 {
			amount, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			newCommand := Command{"s", amount, 0}
			cmds = append(cmds, newCommand)
			continue
		}

		matches = exchange.FindStringSubmatch(cmdStr)
		if len(matches) == 3 {
			pos1, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pos2, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			newCommand := Command{"x", pos1, pos2}
			cmds = append(cmds, newCommand)
			continue
		}

		matches = partner.FindStringSubmatch(cmdStr)
		if len(matches) == 3 {
			newCommand := Command{"p",
				asciiToInt64(matches[1]),
				asciiToInt64(matches[2])}
			cmds = append(cmds, newCommand)
			continue
		}
	}

	//
	// Run all of the commands
	//

	// NOTE: enable this to test input0.txt
	//programOrder := "abcde"

	// NOTE: enable this to test input1.txt
	programOrder := "abcdefghijklmnop"

	for _, cmd := range cmds {

		if cmd.Action == "s" {
			programOrder = conductSpin(programOrder, cmd.First)
		} else if cmd.Action == "x" {
			programOrder = conductExchange(programOrder,
				cmd.First, cmd.Second)
		} else if cmd.Action == "p" {
			programOrder = conductPartner(programOrder,
				int64ToAscii(cmd.First),
				int64ToAscii(cmd.Second))
		}
	}

	fmt.Println("Part 1:", programOrder)

	oneBillion := 1000000000

	// NOTE: enable this to test input0.txt
	//programOrder = "abcde"

	// NOTE: enable this to test input1.txt
	programOrder = "abcdefghijklmnop"

	for j := 0; j < oneBillion; j++ {

		for _, cmd := range cmds {

			if cmd.Action == "s" {
				programOrder = conductSpin(programOrder, cmd.First)
			} else if cmd.Action == "x" {
				programOrder = conductExchange(programOrder,
					cmd.First, cmd.Second)
			} else if cmd.Action == "p" {
				programOrder = conductPartner(programOrder,
					int64ToAscii(cmd.First),
					int64ToAscii(cmd.Second))
			}
		}
	}

	fmt.Println("Part 2:", programOrder)
}

func conductSpin(order string, amount int64) string {

	length := int64(len(order))
	amount = amount % length

	if amount == 0 {
		return order
	}

	newOrder := ""

	start := length - amount
	end := start - 1

	for i := start; i != end; i = (i + 1) % length {
		newOrder += string(order[i])
	}
	newOrder += string(order[end])

	return newOrder
}

func conductExchange(order string, pos1 int64, pos2 int64) string {

	var i int64 = 0
	var length int64 = int64(len(order))

	newOrder := ""

	for i = 0; i < length; i++ {

		char := ""

		if i == pos1 {
			char = string(order[pos2])
		} else if i == pos2 {
			char = string(order[pos1])
		} else {
			char = string(order[i])
		}

		newOrder += char
	}

	return newOrder
}

func conductPartner(order string, program1 string, program2 string) string {

	if len(program1) != 1 || len(program2) != 1 {
		fmt.Println("invalid program name")
		os.Exit(1)
	}

	var pos1 int64 = 0
	var pos2 int64 = 0

	for i := 0; i < len(order); i++ {

		char := string(order[i])

		if char == program1 {
			pos1 = int64(i)
		}

		if char == program2 {
			pos2 = int64(i)
		}
	}

	return conductExchange(order, pos1, pos2)
}

func asciiToInt64(char string) int64 {

	switch char {
	case "a":
		return 61
	case "b":
		return 62
	case "c":
		return 63
	case "d":
		return 64
	case "e":
		return 65
	case "f":
		return 66
	case "g":
		return 67
	case "h":
		return 68
	case "i":
		return 69
	case "j":
		return 70
	case "k":
		return 71
	case "l":
		return 72
	case "m":
		return 73
	case "n":
		return 74
	case "o":
		return 75
	case "p":
		return 76
	}

	return 0
}

func int64ToAscii(char int64) string {

	switch char {

	case 61:
		return "a"
	case 62:
		return "b"
	case 63:
		return "c"
	case 64:
		return "d"
	case 65:
		return "e"
	case 66:
		return "f"
	case 67:
		return "g"
	case 68:
		return "h"
	case 69:
		return "i"
	case 70:
		return "j"
	case 71:
		return "k"
	case 72:
		return "l"
	case 73:
		return "m"
	case 74:
		return "n"
	case 75:
		return "o"
	case 76:
		return "p"
	}

	return ""
}
