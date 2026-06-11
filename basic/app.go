package main

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
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
	fmt.Printf("\n My name is %s Al-Murtadha %q", "Fatima", "Go") // My name is Fatima Al-Murtadha "Go"

	// Padding
	fmt.Printf("\n My name is %9s Al-Murtadha %-9q end", "Fatima", "Go")

	





}
