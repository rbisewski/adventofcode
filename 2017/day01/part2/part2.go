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

	compare := ""

	var value int64 = 0

	for {
		currentChar := string(reverseCaptcha[pos])

		compareInt := (pos + int(end/2)) % end
		compare = string(reverseCaptcha[compareInt])

		fmt.Println(currentChar, " ", compare)
		if currentChar == compare {
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
