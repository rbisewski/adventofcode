package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	x, y int
}

const (
	north int = iota
	east
	south
	west
)

const (
	clean int = iota
	weakened
	infected
	flagged
)

type position [2]int
type direction position

type virus struct {
	p position
	d direction
}

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

	lines := strings.Split(contents, "\n")

	if len(lines) == 0 {
		os.Exit(1)
	}

	//
	// Part 1
	//

	n := len(lines)

	infectionMap := make(map[position]bool)
	for i := range lines {
		for j := range lines {
			if lines[i][j] == '#' {
				infectionMap[[2]int{i, j}] = true
			}
		}
	}

	infected := 0
	virus := virus{[2]int{n / 2, n / 2}, [2]int{-1, 0}}

	for i := 0; i < 10000; i++ {
		if infectionMap[virus.p] {
			virus.d = mirrorRight[virus.d]
			delete(infectionMap, virus.p)
		} else {
			virus.d = mirrorLeft[virus.d]
			infectionMap[virus.p] = true
			infected++
		}
		virus.p[0] += virus.d[0]
		virus.p[1] += virus.d[1]
	}

	fmt.Println("Part 1:", infected)

	//
	// Part 2
	//

	infections := make(map[node]int, 0)

	width, height := len(lines[0]), len(lines)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if string(lines[i][j]) == "#" {
				n := node{j - width/2, i - height/2}
				infections[n] = infected
			}
		}
	}

	icnt, pos, dir := 0, node{0, 0}, north

	for i := 0; i < 10000000; i++ {
		switch infections[pos] {
		case clean:
			dir = (dir + 3) % 4
			infections[pos] = weakened
		case weakened:
			infections[pos] = infected
			icnt++
		case infected:
			dir = (dir + 1) % 4
			infections[pos] = flagged
		case flagged:
			dir = (dir + 2) % 4
			infections[pos] = clean
		}

		switch dir {
		case north:
			pos = node{pos.x, pos.y - 1}
		case east:
			pos = node{pos.x + 1, pos.y}
		case south:
			pos = node{pos.x, pos.y + 1}
		case west:
			pos = node{pos.x - 1, pos.y}
		}
	}

	fmt.Println("Part 2:", icnt)
}

var mirrorRight map[[2]int][2]int = map[[2]int][2]int{
	[2]int{0, 1}:  [2]int{1, 0},
	[2]int{1, 0}:  [2]int{0, -1},
	[2]int{0, -1}: [2]int{-1, 0},
	[2]int{-1, 0}: [2]int{0, 1},
}

var mirrorLeft map[[2]int][2]int = map[[2]int][2]int{
	[2]int{0, 1}:  [2]int{-1, 0},
	[2]int{-1, 0}: [2]int{0, -1},
	[2]int{1, 0}:  [2]int{0, 1},
	[2]int{0, -1}: [2]int{1, 0},
}
