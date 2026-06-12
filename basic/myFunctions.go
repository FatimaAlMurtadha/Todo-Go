package main

import (
	"fmt"
)

func myFunctions(){
	add(2, 3)
	count, count1 := subAndMult(3, 2)
	fmt.Println("Numbers 3 and 2: ")
	fmt.Println("The subtraction is: ",count)
	fmt.Println("The multiplication is: ", count1)

	// call a function inside a variable 
	a := add
	a(1,12)

	// create a function inside a variable
	b := func(blablabla int) int {
		fmt.Println("On run", blablabla)

		return blablabla * -3
	}(22)
	fmt.Println(b)

	test2(addTen) // 17
	test2(mulTen) // 70

	returnFunc("Fatima")() // Welcome Fatima


}

// void function
func add(x, s int){
		fmt.Println("The sum is: ", x + s)
	}

// int OR more
// defer - executed at the end of the function even if it is written first
func subAndMult(x1 , s1 int) (int, int) {
	z := x1 - s1
	c := x1 * s1

	return z, c
}

// --------------------------------------
// --------------------------------------

func test2 (myFunc func(int) int){
	fmt.Println(myFunc(7))
} 
func addTen(number int) int {
	return number + 10
}
func mulTen(number int) int {
	return number * 10
}

// --------------------------------------
// --------------------------------------

func returnFunc(x string) func(){
	return func(){
		fmt.Println("Welcome ", x)
	}
}