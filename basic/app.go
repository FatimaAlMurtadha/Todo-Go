package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, Fatima!")

	var myName string = "Hello from the first string"
	fmt.Println(myName)

	var myNumber int = 33
	fmt.Println("Fatima is ", myNumber, " years old")

	// unit8 - unit16 - unit32 - unit64  -- Unsigned bit
	// int8  - int16  - int32  - int64   -- Signed   bit

	// float32 - float64
	// complex64 - complex128

	// byte = unit8
	// rune = int32
	// uint = 32 or 64 bits
	// int == same size as unit
	// uIntptr

	var number uint8 = 123
	number = 200
	fmt.Println(number) // 200

	// Implicit VS Explicit
	fmt.Printf("%T ", number) // int

	mySecondNumber := 400.17
	fmt.Printf("%T ", mySecondNumber) // float

	// default value int: 0 - bool: false - string: empty string - float: 0
	var check bool
	fmt.Println(check) // false

	// --------------------------------------------------------------------
	//---------------------------------------------------------------------
	
	// fmt printing
	// placeHolder %...
	fmt.Printf("Hello from GoLang the number %v has type %T", 10, 10)
	fmt.Println()

	// text inside variables
	var a = fmt.Sprintf("Hello from GoLang2 the number %v has type %T %t", 10, 10, true)
	fmt.Println(a)

	// converting
	fmt.Printf("Number binary 533 = %b", 533 ) // 1000010101
	fmt.Println()
	fmt.Printf("Number binary 533 = %x", 53393) // d091
	fmt.Printf("\nMy name is %s Al-Murtadha %q", "Fatima", "Go") // My name is Fatima Al-Murtadha "Go"

	// Padding
	fmt.Printf("\nMy name is %9s Al-Murtadha %-9q end\n", "Fatima", "Go")

	// Input
	scanner := bufio.NewScanner(os.Stdin) // A scanner object = Java

	fmt.Println("Enter your name:")
	scanner.Scan()
	// input default is string 
	userInput := scanner.Text()
	fmt.Printf("Your name is %q ", userInput)

	// Convert
	fmt.Println("\nEnter a number: ")
	scanner.Scan()
	numberInput ,_ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("The sum is %d", numberInput + 5)

	// ------------------------------------------------------
	// ------------------------------------------------------
	// Arithmetic operations

	var num1 int = 9
	var num2 int = 20
	answer := num1 + num2
	fmt.Printf("\nYour answer is: %d", answer)

	var num3 float32 = 9
	var num4 float32 = 5
	answer2 := num3 / num4
	fmt.Printf("\nThe division: %g", answer2) // 1.8

	// --------------------------------------------
	// --------------------------------------------
	// Boolean expressions - logical operators
	// < > <= >= == !=

	x := 5
	y := 8
	val := x > y
	fmt.Printf("\n%t", val) // false

	w1 := "fatima"
	w2 := "Fatima"
	fmt.Printf("\n%t", w1 == w2) // false
	w2 = "fatima"
	fmt.Printf("\n%t", w1 == w2) // true

	// ---------------------------------------------
	// ---------------------------------------------
	// Chained Conditionals

	val2 := 9 > 4 && 9 > 11
	fmt.Printf("\nFirst one: %t", val2) // false
	
	val2 = 9 > 4 || 9 > 11
	fmt.Printf("\nSecond one: %t", val2) // true

	val2 = true || false && false // true
	fmt.Printf("\nThird one: %t", val2) // true

	val2 = (true || false) && false 
	fmt.Printf("\nFourth one: %t", val2) // false

	val2 = (true || false) && !false 
	fmt.Printf("\nFifth one: %t", val2) // true

	val3 := val2 || false
	fmt.Printf("\nTow variables: %t", val3) // true

}
