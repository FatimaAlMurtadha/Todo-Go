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

	// call 

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

