package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// init reader object
var reader = bufio.NewReader(os.Stdin)

func main() {
	//calculator for add subtract divide and multiply 2 values
	// value
	value1 := getValue("Value 1")
	value2 := getValue("Value 2")
	// results var
	var result float64
	// calculate based on operator selection
	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")

	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n\n", result)
}

// Get Input Values  int Value to calculate
func getValue(prompt string) float64 {
	fmt.Printf("%v: ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Println("Select an operation (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
