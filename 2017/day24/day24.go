package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	file = ""
)

type Component struct {
	Start int
	End   int
	InUse bool
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

	components := make([]Component, 0)

	re := regexp.MustCompile("([0-9]+)/([0-9]+)")

	for _, l := range lines {

		line := strings.TrimSpace(l)

		matches := re.FindAllStringSubmatch(line, -1)

		if len(matches) < 1 || len(matches[0]) != 3 {
			fmt.Println("Invalid matched set:", matches)
			os.Exit(1)
		}

		start, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			fmt.Println("Improper integer:", matches[0][1])
			os.Exit(1)
		}

		end, err := strconv.ParseInt(matches[0][2], 10, 64)
		if err != nil {
			fmt.Println("Improper integer:", matches[0][2])
			os.Exit(1)
		}

		c := Component{int(start), int(end), false}

		components = append(components, c)
	}

	strongestBridge := 0
	longestBridge := 0
	greatestLength := 0
	cycles := 0

	currentEnd := 0

	zerothComponent := Component{}
	isFirstComponent := true

	bridge := make([]Component, 0)

	for {
		if cycles == 10000000 {
			break
		}

		if isFirstComponent {

			isFirstComponent = false
			components = randomizeOrder(components)

			components, zerothComponent =
				obtainZeroStart(components)

			if zerothComponent.Start == -1 ||
				zerothComponent.End == -1 {

				fmt.Println("Array contains no zero ports!")
				os.Exit(0)
			}

			bridge = append(bridge, zerothComponent)

			if zerothComponent.Start != 0 {
				currentEnd = zerothComponent.Start
			} else {
				currentEnd = zerothComponent.End
			}

			continue

		}

		components, c, openEnd := nextPiece(components, currentEnd)

		// if this reached the end
		if openEnd == -1 {

			newStrongestBridge := sumBridge(bridge)
			newLongestBridge := sumBridge(bridge)
			newGreatestLength := len(bridge)

			if strongestBridge < newStrongestBridge {
				strongestBridge = newStrongestBridge
			}

			if newGreatestLength > greatestLength {

				longestBridge = newLongestBridge
				greatestLength = newGreatestLength

			} else if newGreatestLength == greatestLength &&
				newLongestBridge > longestBridge {

				longestBridge = newLongestBridge
			}

			isFirstComponent = true
			components = resetComponents(components)
			bridge = make([]Component, 0)
			cycles++

			continue
		}

		// else there are still bridge pieces this can use
		currentEnd = openEnd
		bridge = append(bridge, c)
	}

	fmt.Println("Strongest Overall Bridge?", strongestBridge)
	fmt.Println("Strongest-Longest Bridge?", longestBridge, "with a length of",
		greatestLength)
}

func randomizeOrder(array []Component) []Component {

	newArray := make([]Component, 0)

	listOfNumbers := rand.Perm(len(array))

	for _, num := range listOfNumbers {
		newArray = append(newArray, array[num])
	}

	return newArray
}

func obtainZeroStart(array []Component) ([]Component, Component) {

	huskComponent := Component{-1, -1, true}

	for i, element := range array {

		if element.Start == 0 {
			array[i].InUse = true
			return array, array[i]
		}

		if element.End == 0 {
			array[i].InUse = true
			return array, array[i]
		}
	}

	return array, huskComponent
}

func sumBridge(array []Component) int {

	strength := 0

	for _, element := range array {
		strength += element.Start + element.End
	}

	return strength
}

func nextPiece(array []Component, end int) ([]Component, Component, int) {

	huskComponent := Component{-1, -1, true}

	for i, element := range array {

		if element.InUse {
			continue
		}

		if element.Start == end {
			array[i].InUse = true
			return array, array[i], element.End
		}

		if element.End == end {
			array[i].InUse = true
			return array, array[i], element.Start
		}
	}

	// no usable piece was found
	return array, huskComponent, -1
}

func resetComponents(array []Component) []Component {

	for i, _ := range array {
		array[i].InUse = false
	}

	return array
}
