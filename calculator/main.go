package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	str1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(str1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	str2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	float3 := float1 + float2
	float3 = math.Round(float3*100) / 100
	fmt.Printf("The sum of %v and %v is %v", float1, float2, float3)
}
