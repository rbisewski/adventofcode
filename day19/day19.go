package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	lines := strings.Split(contents, "\n")

	if len(lines) == 0 {
		os.Exit(1)
	}

	grid := make([][]string, 0)

	for _, line := range lines {

		if line == "" {
			continue
		}

		row := make([]string, 0)

		for i := 0; i < len(line); i++ {
			char := string(line[i])
			row = append(row, char)
		}

		grid = append(grid, row)
	}

	height := len(grid)
	length := len(grid[0])

	x := 0
	y := 0
	tile := ""
	direction := "south"

	letters := ""
	steps := 0

	// search for the starting point
	for i, char := range grid[0] {
		if char == "|" {
			x = i
			tile = "|"
		}
	}

	// explore thru the maze
	for tile != " " {

		tile = grid[y][x]

		if tile == "+" {

			if direction == "north" || direction == "south" {

				if x-1 >= 0 && grid[y][x-1] == " " {
					direction = "east"

				} else if x+1 < length && grid[y][x+1] == " " {
					direction = "west"
				}

			} else if direction == "east" || direction == "west" {

				if y-1 >= 0 && grid[y-1][x] == " " {
					direction = "south"

				} else if y+1 < height && grid[y+1][x] == " " {
					direction = "north"
				}
			}

		} else if tile != "-" && tile != "|" && tile != "+" {
			letters += tile
		}

		switch direction {
		case "north":
			y--
		case "east":
			x++
		case "south":
			y++
		case "west":
			x--
		}

		if tile != " " {
			steps++
		}
	}

	fmt.Println("Part 1:", letters)
	fmt.Println("Part 2:", steps)
}
