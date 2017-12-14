package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	file = ""
)

type Layer struct {
	Level     int64
	Pos       int64
	Length    int64
	Increment bool
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

	if len(lines) < 1 {
		os.Exit(0)
	}

	layers := make([]Layer, 0)

	var fullLength int64 = 0
	for _, l := range lines {
		nums := strings.Split(l, ": ")
		if len(nums) != 2 {
			continue
		}
		index, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		value, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			os.Exit(1)
		}

		if index > fullLength {
			fullLength = index
		}

		newLayer := Layer{index, 0, value, true}
		layers = append(layers, newLayer)
	}

	var severity int64
	var picos int64

	// for every picosecond
	for picos = 0; picos <= fullLength; picos++ {

		layer, err := obtainLayerAtLevel(layers, picos)

		if err == nil && layer.Pos == 0 {
			severity += (picos * layer.Length)
		}

		// increment or decrement pos
		for i, _ := range layers {

			if layers[i].Pos == 0 {
				layers[i].Increment = true
			} else if layers[i].Pos == layers[i].Length-1 {
				layers[i].Increment = false
			}

			if layers[i].Increment {
				layers[i].Pos++
			} else {
				layers[i].Pos--
			}
		}
	}

	fmt.Println(severity)
}

func obtainLayerAtLevel(layers []Layer, level int64) (Layer, error) {

	if len(layers) == 0 || level < 0 {
		return Layer{}, fmt.Errorf("invalid input")
	}

	for _, l := range layers {
		if l.Level == level {
			return l, nil
		}
	}

	return Layer{}, fmt.Errorf("layer not found")
}
