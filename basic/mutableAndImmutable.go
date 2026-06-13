package main

import (
	"fmt"
)

func MutableAndImmutable(){
	
// --------------------------------------
// --------------------------------------
// Mutable & Immutable data types

var x [] int = [] int { 3, 4, 5}
y := x // 
y[1] = 50 // both changed - mutable because the reference is the same for both

fmt.Println(x , y)
fmt.Println()

var x2 map[string] int = map[string]int {"a": 5 }
y2 := x2
y2["b"] = 60
x2["c"] = 80
fmt.Println("First one:  ",x2, "\nSecond one: ", y2)

var x3 [2]int = [2]int {3,4}
y3 := x3
y3[0] = 90 // only y3 has changed - they have n't the same reference type - The value type can't be changed

fmt.Println("First one:  ", x3, "\nSecond one: ", y3)

var x4 [] int = []int {3,4,5}

fmt.Println("Before: ",x4)
updateSlice(x4)
fmt.Println("After:  ",x4)

// Slice + Map + Variables == reference type == MUTABLE
// Arrays == value type == IMMUTABLE
// stack - managed heap

}

func updateSlice(slice [] int){
	slice[0] =55
}