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

type Particle struct {
	ID                int
	P                 Vertex
	V                 Vertex
	A                 Vertex
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

	particles := make(map[int]Particle, 0)
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

		position := Vertex{int(positionX), int(positionY),
			int(positionZ)}
		velocity := Vertex{int(velocityX), int(velocityY),
			int(velocityZ)}
		acceleration := Vertex{int(accelerationX), int(accelerationY),
			int(accelerationZ)}

		particles[p] = Particle{p, position, velocity,
			acceleration, 0}

		p++
	}

	closestParticle := -1
	counter := 0
	max := 2000000

	for counter < max {

		id := -1
		closestDistance := -1

		for key, ptc := range particles {

			thisParticle := particles[key]

			sum := int(math.Abs(float64(ptc.P.X)) +
				math.Abs(float64(ptc.P.Y)) +
				math.Abs(float64(ptc.P.Z)))

			thisParticle.ManhattanDistance = sum

			thisParticle.V.X += thisParticle.A.X
			thisParticle.V.Y += thisParticle.A.Y
			thisParticle.V.Z += thisParticle.A.Z

			thisParticle.P.X += thisParticle.V.X
			thisParticle.P.Y += thisParticle.V.Y
			thisParticle.P.Z += thisParticle.V.Z

			particles[key] = thisParticle

			if (id == -1) ||
				(thisParticle.ManhattanDistance < closestDistance) {

				id = thisParticle.ID
				closestDistance = thisParticle.ManhattanDistance
			}
		}

		if closestParticle == -1 {
			closestParticle = id
			counter = 1
		} else if id == closestParticle {
			counter++
		} else {
			closestParticle = -1
			counter = 0
		}
	}

	fmt.Println("Part 1:", closestParticle)
}
