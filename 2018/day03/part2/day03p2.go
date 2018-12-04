package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	flength = 1000
)

var (
	inputFile    = ""
	fabricMatrix [flength][flength]int
)

type Claim struct {
	id     int
	startX int
	startY int
	endX   int
	endY   int
}

func init() {
	flag.StringVar(&inputFile, "inputFile", "",
		"Enter the input file location.")
}

func main() {

	flag.Parse()

	if inputFile == "" {
		fmt.Println("No file was specified.")
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		os.Exit(1)
	}

	fileContents := string(bytes)

	array := strings.Split(fileContents, "\n")

	blankElementIndex := len(array) - 1

	//
	// obtain a list of claims
	//
	listOfClaims := make([]Claim, len(array))
	for i, str := range array {

		if i == blankElementIndex {
			continue
		}

		claim, err := parseLine(str)
		if err != nil {
			fmt.Println("Invalid claim detected!")
			fmt.Println("Terminating program.")
			os.Exit(1)
		}

		listOfClaims[i] = claim
	}

	//
	// increment the fabric matrix
	//
	for _, claim := range listOfClaims {
		incrementFabricSquare(claim)
	}

	//
	// cycle thru the IDs
	//
	for i := 1; i <= flength; i++ {
		if !fabricSquareOverlaps(listOfClaims[i-1]) {
			fmt.Println(i)
		}
	}
}

func parseLine(line string) (Claim, error) {

	if line == "" {
		return Claim{}, fmt.Errorf("Invalid input")
	}

	idAndAttributes := strings.Split(line, " @ ")
	id, _ := strconv.Atoi(strings.TrimPrefix(idAndAttributes[0], "#"))
	attributes := idAndAttributes[1]

	coordsAndDimensions := strings.Split(attributes, ": ")
	coords := coordsAndDimensions[0]
	dimensions := coordsAndDimensions[1]

	startXAndstartY := strings.Split(coords, ",")
	startX, _ := strconv.Atoi(startXAndstartY[0])
	startY, _ := strconv.Atoi(startXAndstartY[1])

	lengthAndHeight := strings.Split(dimensions, "x")
	length, _ := strconv.Atoi(lengthAndHeight[0])
	height, _ := strconv.Atoi(lengthAndHeight[1])

	newClaim := Claim{}

	newClaim.id = id
	newClaim.startX = startX
	newClaim.startY = startY
	newClaim.endX = startX + length
	newClaim.endY = startY + height

	return newClaim, nil
}

func incrementFabricSquare(claim Claim) {

	for i := claim.startX; i < claim.endX; i++ {
		for j := claim.startY; j < claim.endY; j++ {
			fabricMatrix[i][j]++
		}
	}
}

func fabricSquareOverlaps(claim Claim) bool {

	overlaps := false

	for i := claim.startX; i < claim.endX; i++ {
		for j := claim.startY; j < claim.endY; j++ {
			if fabricMatrix[i][j] > 1 {
				overlaps = true
			}
		}
	}

	return overlaps
}
