package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	float1 := getValue("1")
	float2 := getValue("2")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(float1, float2)
	case "-":
		result = substractValues(float1, float2)
	case "*":
		result = multiplyValues(float1, float2)
	case "/":
		result = divideValues(float1, float2)
	default:
		panic(" Operation must be one of these + - * /")
	}

	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n\n", result)

}

func getValue(valueNumber string) float64 {
	fmt.Printf("Value %v: ", valueNumber)
	input, _ := reader.ReadString('\n')
	floatIn, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Value %v must be a number", valueNumber))
	}
	return floatIn
}

func getOperation() string {
	fmt.Print("Select an operation (+ - * /): ")
	opInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(opInput)
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func substractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
