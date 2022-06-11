package main

import (
	"fmt"
	"sort"
)

func main() {
	//define slice [] instead of array [3] (immutable )
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	// append items to slice
	colors = append(colors, "purple")
	fmt.Println(colors)
	// remove items
	// colors[1:len(colors)] delete first item and shift left remainder of slice
	//colors = append(colors[1:len(colors)]) #len(colors) unecessary
	colors = append(colors[1:])
	fmt.Println(colors)
	// remove last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// initialize slice with 5 elements, 5 preallocated
	numbers := make([]int, 5, 5)
	// initialize slice with size 5, omitted preallocation
	//numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	//numbers[5] = 666  err index out of range

	fmt.Println(numbers)
	//append
	numbers = append(numbers, 235)
	fmt.Println(numbers)

	// sorting
	sort.Ints(numbers)
	fmt.Println(numbers)

	// sorting
	colors = append(colors, "yellow")
	colors = append(colors, "velvet", "black", "white")
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)

}
