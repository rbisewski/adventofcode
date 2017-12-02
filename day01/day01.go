package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	reverseCaptcha = ""
)

func init() {
	flag.StringVar(&reverseCaptcha, "captcha", "",
		"Enter the reverse captcha.")
}

func main() {

	flag.Parse()

	if reverseCaptcha == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	pos := 0
	end := len(reverseCaptcha)

	previous := ""

	var value int64 = 0

	for {
		currentChar := string(reverseCaptcha[pos])

		if pos == 0 {
			previous = string(reverseCaptcha[end-1])
		} else {
			previous = string(reverseCaptcha[pos-1])
		}

		fmt.Println(currentChar, " ", previous)
		if currentChar == previous {
			intter, err := strconv.ParseInt(currentChar, 10, 64)
			if err != nil {
				break
			}
			value += intter
		}

		pos++
		if pos == end {
			break
		}
	}

	fmt.Println(value)
}
