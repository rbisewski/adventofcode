package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	passphrase_list = ""
)

func init() {
	flag.StringVar(&passphrase_list, "file", "",
		"Enter the passphrase filepath.")
}

func main() {

	flag.Parse()

	if passphrase_list == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(passphrase_list)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\n")

	var numberOfValidPassphrases int64 = 0

	for _, line := range lines {

		elements := strings.Split(line, " ")
		if len(elements) == 0 {
			continue
		} else if len(elements) == 1 && elements[0] == "" {
			continue
		}

		useLine := true
		hashnik := make(map[string]int)
		for _, e := range elements {

			if e == "" {
				continue
			}

			_, alreadyExists := hashnik[e]
			if alreadyExists {
				useLine = false
				break
			}

			hashnik[e] = 1
		}

		if useLine {
			numberOfValidPassphrases++
		}
	}

	fmt.Println(numberOfValidPassphrases)
}
