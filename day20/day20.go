package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	file = ""
)

func init() {
	flag.StringVar(&file, "file", "",
		"Enter the file filepath.")
}

type Vertex struct {
	X int
	Y int
	Z int
}

type Instruction struct {
	Particle int
	P        Vertex
	V        Vertex
	A        Vertex
}

type Particle struct {
	ID                int
	ManhattanDistance int
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

	re := regexp.MustCompile("p=<(-?\\d+),(-?\\d+),(-?\\d+)>, v=<(-?\\d+),(-?\\d+),(-?\\d+)>, a=<(-?\\d+),(-?\\d+),(-?\\d+)>")

	// assemble instructions
	instructions := make([]Instruction, 0)
	p := 0
	for _, line := range lines {

		matches := re.FindStringSubmatch(line)

		positionX, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		positionY, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		positionZ, err := strconv.ParseInt(matches[3], 10, 64)
		if err != nil {
			os.Exit(1)
		}

		velocityX, err := strconv.ParseInt(matches[4], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		velocityY, err := strconv.ParseInt(matches[5], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		velocityZ, err := strconv.ParseInt(matches[6], 10, 64)
		if err != nil {
			os.Exit(1)
		}

		accelerationX, err := strconv.ParseInt(matches[7], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		accelerationY, err := strconv.ParseInt(matches[8], 10, 64)
		if err != nil {
			os.Exit(1)
		}
		accelerationZ, err := strconv.ParseInt(matches[9], 10, 64)
		if err != nil {
			os.Exit(1)
		}

		position := Vertex{int(positionX), int(positionY), int(positionZ)}
		velocity := Vertex{int(velocityX), int(velocityY), int(velocityZ)}
		acceleration := Vertex{int(accelerationX), int(accelerationY), int(accelerationZ)}

		instruct := Instruction{p, position, velocity, acceleration}

		instructions = append(instructions, instruct)

		p = (p + 1) % 2
	}

	// assemble particles
	particles := make(map[int]Particle)
	particles[0] = Particle{0, 0}
	particles[1] = Particle{1, 0}

	// compute particle positions
	for _, instr := range instructions {

		sum := int(math.Abs(float64(instr.P.X)) + math.Abs(float64(instr.P.Y)) + math.Abs(float64(instr.P.Z)))

		thisParticle := particles[instr.Particle]

		thisParticle.ManhattanDistance += sum

		particles[instr.Particle] = thisParticle
	}

	// figure out closest particle over time
	closestParticle := 0
	nearestValue := 1000000000000000
	for _, p := range particles {

		if p.ManhattanDistance < nearestValue {
			closestParticle = p.ID
			nearestValue = p.ManhattanDistance
		}
	}

	fmt.Println("Part 1:", closestParticle)
}
