package main

import (
	"fmt"
)

func main() {
	// ------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------
	// Arrays

	var arr [5] int 
	arr[0] = 100
	arr[0] = 6
	fmt.Println(arr)
	fmt.Println(arr[3])
	fmt.Println(len(arr))

	arr1 := [3] int {7, 8,9}
	fmt.Println(arr1)
	fmt.Println(len(arr1))

	// we can use for loop to iterate over the array as usual
	// { {1,2}, {3,4} , {5,6}} 2D array
	arr2D := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[1][1]) // 4 == Python

}