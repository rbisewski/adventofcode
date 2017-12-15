package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	file = ""
)

type MatrixElement struct {
	IsUsed bool
	Group  int64
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

	fullOutput := ""
	matrix := make([][]MatrixElement, 0)

	for i := 0; i < 128; i++ {

		key := contents + "-" + strconv.FormatInt(int64(i), 10)

		hash, err := knotHash(key)
		if err != nil {
			os.Exit(1)
		}

		binaryHash, err := binarify(hash)
		if err != nil {
			os.Exit(1)
		}

		gridLine := convertLineToGridLine(binaryHash)
		matrix = append(matrix, gridLine)

		fullOutput += binaryHash + "\n"
	}

	re := regexp.MustCompile("1")

	matches := re.FindAllString(fullOutput, -1)

	fmt.Println("Part 1: ", len(matches))

	//grpMatrix, _ := assignGroups(matrix)
	_, count := assignGroups(matrix)

	//groups := make(map[int64]int)
	//for _, a := range grpMatrix {

	//	for _, b := range a {
	//		groups[b.Group]++
	//	}
	//}

	fmt.Println("Part 2: ", count)
}

func knotHash(key string) (string, error) {

	if key == "" {
		return "", fmt.Errorf("invalid input")
	}

	bytes := []byte(key)

	addendum := []byte{17, 31, 73, 47, 23}

	for _, a := range addendum {
		bytes = append(bytes, a)
	}

	// generate list of numbers from 0 to 255
	list := make([]int64, 0)
	var counter int64 = 0
	for counter < 256 {
		list = append(list, counter)
		counter++
	}

	// convert bytes to int lengths
	lengths := make([]int64, 0)
	for _, b := range bytes {
		newLength := int64(b)
		lengths = append(lengths, newLength)
	}

	var pos int64 = 0
	var skip int64 = 0

	for i := 0; i < 64; i++ {
		list, pos, skip = executeRound(list, lengths, pos, skip)
	}

	hashnik, err := densifyHash(list)
	if err != nil {
		os.Exit(1)
	}

	finalHash := ""
	for _, decimal := range hashnik {
		hex := fmt.Sprintf("%x", decimal)
		if len(hex) == 1 {
			hex = "0" + hex
		}
		finalHash += hex
	}

	return finalHash, nil
}

func flipElements(list []int64, pos int64, length int64) []int64 {

	if length < 2 {
		return list
	}

	start := pos
	end := (pos + length) % int64(len(list))
	pos2 := start

	// obtain slice
	slice := make([]int64, 0)
	for i := start; i != end; i = (i + 1) % int64(len(list)) {
		slice = append(slice, list[i])
	}

	// inverse swap
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// re-insert
	for _, element := range slice {
		list[pos2] = element
		pos2 = (pos2 + 1) % int64(len(list))
	}

	return list
}

func executeRound(list []int64, lengths []int64, pos int64, skip int64) ([]int64,
	int64, int64) {

	var listSize int64 = int64(len(list))

	for _, l := range lengths {

		list = flipElements(list, pos, l)
		pos = (pos + l + skip) % listSize
		skip++
	}

	return list, pos, skip
}

func densifyHash(list []int64) ([]int64, error) {

	if len(list) < 1 {
		return nil, fmt.Errorf("invalid or empty list")
	}

	remainder := len(list) % 16

	if remainder != 0 {
		return nil, fmt.Errorf("list is not a multiple of 16")
	}

	newList := make([]int64, 0)

	var element int64 = 0
	count := 1
	for _, l := range list {

		if count == 1 {
			element = l
			count++
			continue
		}

		element ^= l

		if count == 16 {
			count = 1
			newList = append(newList, element)
			continue
		}

		count++
	}

	return newList, nil
}

func convertOneHexToBinary(hex string) (string, error) {

	if hex == "" || len(hex) != 1 {
		return "", fmt.Errorf("invalid input")
	}

	value := ""

	switch hex {
	case "0":
		value = "0000"
	case "1":
		value = "0001"
	case "2":
		value = "0010"
	case "3":
		value = "0011"
	case "4":
		value = "0100"
	case "5":
		value = "0101"
	case "6":
		value = "0110"
	case "7":
		value = "0111"
	case "8":
		value = "1000"
	case "9":
		value = "1001"
	case "a":
		value = "1010"
	case "b":
		value = "1011"
	case "c":
		value = "1100"
	case "d":
		value = "1101"
	case "e":
		value = "1110"
	case "f":
		value = "1111"
	}

	return value, nil
}

func binarify(hex string) (string, error) {

	if hex == "" {
		return "", fmt.Errorf("invalid input")
	}

	result := ""

	for i := 0; i < len(hex); i++ {

		char := string(hex[i])
		out, err := convertOneHexToBinary(char)
		if err != nil {
			return "", err
		}

		result += out
	}

	return result, nil
}

func convertLineToGridLine(line string) []MatrixElement {

	gridLine := make([]MatrixElement, 0)

	for i := 0; i < len(line); i++ {

		char := string(line[i])

		used := true
		if char == "0" {
			used = false
		}

		m := MatrixElement{used, 0}

		gridLine = append(gridLine, m)
	}

	return gridLine
}

func assignGroups(matrix [][]MatrixElement) ([][]MatrixElement, int64) {

	var numberOfGroups int64 = 0

	for x, i := range matrix {

		for y, j := range i {

			// skip if free
			if !j.IsUsed {
				continue
			}

			// skip if already has a group
			if j.Group > 0 {
				continue
			}

			makeNewGroup := true

			// (x+1, y)
			if x+1 < len(i) {
				if matrix[x+1][y].Group > 0 {
					matrix[x][y].Group = matrix[x+1][y].Group
					makeNewGroup = false
				}

			}

			// (x-1, y)
			if x-1 >= 0 {
				if matrix[x-1][y].Group > 0 {
					matrix[x][y].Group = matrix[x-1][y].Group
					makeNewGroup = false
				}

			}

			// (x, y+1)
			if y+1 < len(matrix) {
				if matrix[x][y+1].Group > 0 {
					matrix[x][y].Group = matrix[x][y+1].Group
					makeNewGroup = false
				}

			}

			// (x, y-1)
			if y-1 >= 0 {
				if matrix[x][y-1].Group > 0 {
					matrix[x][y].Group = matrix[x][y-1].Group
					makeNewGroup = false
				}
			}

			// otherwise this is a new group, so give it a new
			// number and increment the count
			if makeNewGroup {
				numberOfGroups++
				matrix[x][y].Group = numberOfGroups
			}

			thisGroup := matrix[x][y].Group

			checkAndSet(matrix, thisGroup, x, y, len(i), len(matrix))
		}
	}

	return matrix, numberOfGroups
}

func checkAndSet(matrix [][]MatrixElement, thisGroup int64, x int, y int,
	maxX int, maxY int) {

	// (x+1, y)
	if x+1 < maxX {
		if matrix[x+1][y].IsUsed && matrix[x+1][y].Group != thisGroup {

			matrix[x+1][y].Group = thisGroup

			checkAndSet(matrix, thisGroup, x+1,
				y, maxX, maxY)
		}
	}

	// (x-1, y)
	if x-1 >= 0 {
		if matrix[x-1][y].IsUsed && matrix[x-1][y].Group != thisGroup {

			matrix[x-1][y].Group = thisGroup

			checkAndSet(matrix, thisGroup, x-1,
				y, maxX, maxY)
		}

	}

	// (x, y+1)
	if y+1 < maxY {
		if matrix[x][y+1].IsUsed && matrix[x][y+1].Group != thisGroup {

			matrix[x][y+1].Group = thisGroup

			checkAndSet(matrix, thisGroup, x,
				y+1, maxX, maxY)
		}

	}

	// (x, y-1)
	if y-1 >= 0 {
		if matrix[x][y-1].IsUsed && matrix[x][y-1].Group != thisGroup {

			matrix[x][y-1].Group = thisGroup

			checkAndSet(matrix, thisGroup,
				x, y-1, maxX, maxY)
		}
	}
}
