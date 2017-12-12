package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

var (
	file = ""
)

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

	directions := strings.Split(contents, ",")

	x := 0.0
	y := 0.0
	z := 0.0

	distances := make([]float64, 0)

	for _, d := range directions {

		if d == "n" {
			y++
			z--
		} else if d == "s" {
			y--
			z++
		} else if d == "ne" {
			x++
			z--
		} else if d == "sw" {
			x--
			z++
		} else if d == "nw" {
			x--
			y++
		} else if d == "se" {
			x++
			y--
		}

		length := (math.Abs(x) + math.Abs(y) + math.Abs(z)) / 2.0

		distances = append(distances, length)
	}

	fmt.Println("Fewest steps?")
	fmt.Println((math.Abs(x) + math.Abs(y) + math.Abs(z)) / 2.0)
	fmt.Println("Furthest away?")
	fmt.Println(max(distances))
}

func max(array []float64) float64 {

	largest := array[0]
	for _, a := range array {
		if a > largest {
			largest = a
		}
	}
	return largest
}
