package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from Go!")
	doSomething()
	sum := addValues(5, 8)
	fmt.Println("the sum is", sum)

	multiSum, multiCount := addAllValues(4, 7, 9, 12)
	fmt.Println("multiSum", multiSum)
	fmt.Println("multiCount", multiCount)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
