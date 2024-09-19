package main

import (
	"fmt"
)

type coordinate [2]int

func main() {
	var locs [5]coordinate

	locs[0] = [2]int{1, 2}
	locs[1] = [...]int{7, 8}
	locs[2] = coordinate{3, 4}
	locs[3][0] = 5
	locs[4][1] = 6

	numSlice := []int{1, 2, 3, 4, 5}
	locsSlice := locs[:4]
	falseSlice := make([]bool, 1, 5)

	fmt.Println(locs)
	fmt.Println(locsSlice, numSlice, falseSlice)

	fmt.Println(
		len(locsSlice),
		cap(locsSlice),
		append(locsSlice, coordinate{7, 7}),
	)

}
