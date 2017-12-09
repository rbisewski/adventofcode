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
	towerList = ""
)

type Node struct {
	Name   string
	Weight int64
	Subs   []string
}

func init() {
	flag.StringVar(&towerList, "file", "",
		"Enter the tower list filepath.")
}

func main() {

	flag.Parse()

	if towerList == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(towerList)

	if err != nil {
		os.Exit(1)
	}

	filestr := string(bytes)

	if filestr == "" {
		fmt.Println(0)
		os.Exit(0)
	}

	lines := strings.Split(filestr, "\n")

	if len(lines) < 1 {
		fmt.Println("lines are empty")
		os.Exit(0)
	}

	bottom, err := determineBottomProgram(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get the weights
	nodes := make([]Node, 0)
	for _, l := range lines {

		if len(l) == 0 {
			continue
		}

		elements := strings.Split(l, " ")

		name := elements[0]
		weight := elements[1]

		weight = strings.Trim(weight, "(")
		weight = strings.Trim(weight, ")")

		weightAsInt, err := strconv.ParseInt(weight, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subs := make([]string, 0)

		for i, e := range elements {
			if i < 2 {
				continue
			}

			if e == "->" {
				continue
			}

			e2 := strings.Trim(e, ",")

			subs = append(subs, e2)
		}

		newNode := Node{name, weightAsInt, subs}
		nodes = append(nodes, newNode)
	}

	// get the bottom node
	bottomNode, err := getNode(bottom, nodes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var requiredWeight int64 = 0

	parent := bottomNode
	for {
		unevenSubnode, offset := checkForUnevenSubnode(parent,
			nodes)
		if unevenSubnode.Name != "" {
			if len(unevenSubnode.Subs) == 0 {
				break
			}
			requiredWeight = unevenSubnode.Weight - offset
			parent = unevenSubnode
		}
	}

	fmt.Println(requiredWeight)
}

func determineBottomProgram(lines []string) (string, error) {

	if len(lines) == 0 {
		return "", fmt.Errorf("invalid line data")
	}

	arrow_lines := make([]string, 0)
	for _, l := range lines {

		if strings.Contains(l, " -> ") {
			arrow_lines = append(arrow_lines, l)
		}
	}

	if len(arrow_lines) < 1 {
		return "", fmt.Errorf("arrow_lines is empty")
	}

	bottomProgram := ""
	for _, l := range arrow_lines {

		elements := strings.Split(l, " ")

		if len(elements) < 1 {
			return "", fmt.Errorf("elements is empty")
		}

		first := elements[0]

		programIsSub := false
		for _, l2 := range arrow_lines {

			if strings.Contains(l2, " "+first) {

				programIsSub = true
				break
			}
		}

		if !programIsSub {
			bottomProgram = first
			break
		}
	}

	return bottomProgram, nil
}

func getNode(name string, listOfNodes []Node) (Node, error) {

	if name == "" {
		return Node{}, fmt.Errorf("invalid input for node name")
	}

	for _, node := range listOfNodes {

		if node.Name == name {
			return node, nil
		}
	}

	return Node{}, fmt.Errorf("node name not found!")
}

func calculateWeight(node Node, nodes []Node) int64 {

	if len(node.Subs) == 0 {
		return node.Weight
	}

	var value int64 = node.Weight
	for _, sub := range node.Subs {
		newNode, _ := getNode(sub, nodes)
		value += calculateWeight(newNode, nodes)
	}
	return value
}

func checkForUnevenSubnode(node Node, nodes []Node) (Node, int64) {

	// get the subnodes of the bottom node
	seconders := make([]Node, 0)
	for _, sub := range node.Subs {
		subnode, err := getNode(sub, nodes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		seconders = append(seconders, subnode)
	}

	secondersValues := make(map[string]int64, 0)
	for _, sub := range seconders {

		value := calculateWeight(sub, nodes)
		secondersValues[sub.Name] = value
	}

	highestName := ""
	var highestValue int64 = 0
	var lowestValue int64 = 0
	for key, value := range secondersValues {

		if highestName == "" {
			highestName = key
			highestValue = value
			lowestValue = value
			continue
		}

		if value > highestValue {
			highestValue = value
			highestName = key
		}

		if value < lowestValue {
			lowestValue = value
		}
	}

	newNode, err := getNode(highestName, nodes)
	if err != nil {
		return Node{}, -1
	}

	return newNode, highestValue - lowestValue
}
