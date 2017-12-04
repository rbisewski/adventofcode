package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type vector struct {
	X int
	Y int
}

var (
	square = 0
)

func init() {
	flag.IntVar(&square, "square", 0,
		"Enter the data square, to be used to calculate steps.")
}

func main() {

	flag.Parse()

	if square < 2 {
		fmt.Println(0)
		os.Exit(0)
	}

	length_of_square := int(math.Ceil(math.Sqrt(float64(square))))

	// enforce squares of odd length
	if length_of_square%2 == 0 {
		length_of_square++
	}

	one_x := (length_of_square - 1) / 2
	one_y := one_x

	cur_x := length_of_square - 1
	cur_y := cur_x

	// largest possible x or y value for the outter rung
	highest := cur_x

	// current ending point for the outer rung
	current := length_of_square * length_of_square

	// direction bools
	traveling_west := true
	traveling_north := false
	traveling_east := false
	traveling_south := false

	for {

		if current == square {
			break
		}

		if traveling_west {
			cur_x--
			if cur_x == 0 {
				traveling_west = false
				traveling_north = true
			}
			current--

		} else if traveling_north {
			cur_y--
			if cur_y == 0 {
				traveling_north = false
				traveling_east = true
			}
			current--

		} else if traveling_east {
			cur_x++
			if cur_x == highest {
				traveling_east = false
				traveling_south = true
			}
			current--

		} else if traveling_south {
			cur_y++
			if cur_y == highest {
				traveling_east = false
				traveling_south = false
			}
			current--

		} else {
			// this should never happen
			break
		}
	}

	distance := math.Abs(float64(cur_x-one_x)) +
		math.Abs(float64(cur_y-one_y))

	fmt.Println(distance)
}
