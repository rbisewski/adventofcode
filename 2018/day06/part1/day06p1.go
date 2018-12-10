package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
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

	contentsAsString := strings.TrimSpace(string(bytes))

	contents := strings.Split(contentsAsString, "\n")

	coords := make([]Coord, len(contents))

	for i, str := range contents {
		values := strings.Split(str, ", ")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		coordinate := Coord{x: x, y: y}
		coords[i] = coordinate
	}

	// create a value that substitutes for infinity
	infinity := 1000

	array := make([][]string, infinity, infinity)

	for i := 0; i < infinity; i++ {
		for j := 0; j < infinity; j++ {

			closestAbsoluteLength := infinity
			closestCoord := ""

			for num, coord := range coords {

				absoluteLength := int(math.Abs(float64(i-coord.x)) + math.Abs(float64(j-coord.y)))

				if absoluteLength == closestAbsoluteLength {
					closestCoord = "."
					break

				} else if absoluteLength < closestAbsoluteLength {
					closestAbsoluteLength = absoluteLength
					closestCoord = strconv.Itoa(num)
				}
			}

			array[i][j] = closestCoord
		}
	}

	fmt.Println(array)
}
