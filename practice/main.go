package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "My String"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var aInteger int = 42
	fmt.Println(aInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)
}
