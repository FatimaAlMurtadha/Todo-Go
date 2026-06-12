package main

import (
	"fmt"
)

func Arrays() {
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
	fmt.Println("2D array: ",arr2D)
	fmt.Println("\nIndex [1][1]: " , arr2D[1][1]) // 4 == Python

	// ------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------
	// Slices
	var x [6] int = [6] int {4, 5, 6, 7, 8, 9}
	var s [] int = x [1:3] // [5, 6]

	fmt.Println("\nThe array is: ",x , "\nThe slice is: ", s)
	fmt.Println("Length of the slice:  ",len(s)) // 2
	fmt.Println("Capacity of the slice: " , cap(s)) // 5

	// slice out of the slice
	fmt.Println(s[1:cap(s)]) // [6 7 8 9]
	b := append(s, 100)
	fmt.Println(b) // [5 , 6 , 100]

	m := make([]int, 5)
	fmt.Println(m) // [0 0 0 0 0]

	// Array range
	var a []int = [] int {1,2,3,4,5,6,7,8, 3}

	/*for i := 0 ; i< len(a); i ++ {
		fmt.Println(a[i])
	}*/

	// i and element == foreach
	// if we don't want the index so we can delete the i and use instead _ as a placeholder

	for f, element := range a{
		fmt.Printf("%d: %d \n", f , element) // indexes with values
	}

	fmt.Println()
	for j, el := range a {
		for y, el1 := range a {
			if el == el1 && j !=y {
				fmt.Println(el)
			}
		}
	}

	// Maps






}