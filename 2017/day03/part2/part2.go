package main

import (
	"flag"
	"fmt"
	"os"
)

type Vector struct {
	X     int
	Y     int
	Value int
}

var (
	square = 0
)

func init() {
	flag.IntVar(&square, "square", 0,
		"Enter the data square input.")
}

func main() {

	flag.Parse()

	if square < 1 {
		fmt.Println(0)
		os.Exit(0)
	}

	cur_x := 0
	cur_y := 0

	current := 1
	min_x := 0
	min_y := 0
	max_x := 0
	max_y := 0

	// direction bools
	traveling_east := true
	traveling_north := false
	traveling_west := false
	traveling_south := false

	vec_array := make([]Vector, 0)
	first := Vector{0, 0, 1}

	vec_array = append(vec_array, first)

	for {
		if current > square {
			break
		}

		if traveling_east {
			cur_x++
			if cur_x > max_x {
				max_x++
				min_x = -1 * max_x
				traveling_east = false
				traveling_north = true
			}

		} else if traveling_north {
			cur_y++
			if cur_y > max_y {
				max_y++
				min_y = -1 * max_y
				traveling_north = false
				traveling_west = true
			}

		} else if traveling_west {
			cur_x--
			if cur_x == min_x {
				traveling_west = false
				traveling_south = true
			}

		} else if traveling_south {
			cur_y--
			if cur_y == min_y {
				traveling_south = false
				traveling_east = true
			}
		}

		current = determineValue(cur_x, cur_y, vec_array)

		new_vec := Vector{cur_x, cur_y, current}
		vec_array = append(vec_array, new_vec)
	}

	fmt.Println(current)
}

func determineValue(x int, y int, vec_array []Vector) int {

	if len(vec_array) < 2 {
		return 1
	}

	value := 0
	for _, vec := range vec_array {

		if (vec.X+1) == x || (vec.X == x) || (vec.X-1 == x) {
			if (vec.Y+1) == y || (vec.Y == y) || (vec.Y-1 == y) {
				value += vec.Value
			}
		}
	}

	return value
}
