package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

			sorted := SortString(e)

			_, alreadyExists := hashnik[sorted]
			if alreadyExists {
				useLine = false
				break
			}

			hashnik[sorted] = 1
		}

		if useLine {
			numberOfValidPassphrases++
		}
	}

	fmt.Println(numberOfValidPassphrases)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
