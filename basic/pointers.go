package main

import (
	"fmt"
)

func MyPointers(){
	x := 8

	// How do I know what's the reference ??
	fmt.Println(&x, "Value: ", x) // 0x11d6d7c380e0 - physical reference


	y := x

	y = 5

	fmt.Println(&x, "Value before the pointer: ", x) // 0x2ce6096c0e0  Value:  8
	fmt.Println(&y, "Value before the pointer: ", y) // 0x2ce6096c0e8  Value:  5
	fmt.Println()

	// Giving the same pointer
	y2 := &x

	fmt.Println(&x, "Value x after the pointer:   ", x)   // 0x2bfc1d7680e0 Value x after the pointer:  8
	fmt.Println(&y2, "Value y2 after the pointer: ", y2) // 0x2bfc1d762068 Value y2 after the pointer:  0x2bfc1d7680e0
	fmt.Println()

	// The value - Change the pointer of y 

	*y2 = 5

	fmt.Println(&x, "Value x after * :  ", x)    // 0x25fe8798e070 Value x after * :   5
	fmt.Println(&y2, "Value y2 after * : ", *y2) // 0x25fe87982040 Value y2 after * :  5

	x2 := "value"
	fmt.Println("Before setValue: ",x2) // Before setValue:  value

	setValue(&x2)
	fmt.Println("After setValue: ",x2) // After setValue:  Done
	fmt.Println()

	fmt.Println("Before setValue2: ",x2) // Before setValue2:  Done

	setValue2(x2)
	fmt.Println("After setValue2: ",x2) // After setValue2:  Done

	x2 = setValue2(x2)
	fmt.Println("After assigning setValue2: ",x2) // After assigning setValue2:  Done-2


	// Summary
	var x3 string = "Hi"
	var pointer *string = &x3 
	fmt.Println("The reference address: ", pointer, "\nMy pointer address: ",&pointer ,"\nThe pointer value: ", *pointer) 

	x4 := 10
	p := &x4 

	fmt.Println()
	fmt.Println("Address x4 = ", &x4) 								   // 0x1f2cf0f8e088
	fmt.Println("Value x4 = ", x4) 									  // 10
	fmt.Print("x4 address  - reference address           = ", p)  	 // 0x1f2cf0f8e088 
	fmt.Println("\nThe box address - pointer address       = ", &p) //  0x1f2cf0f82050 - ITS OWN 
	fmt.Print("x4 value - and pointer value              = ", *p)  // 10

	
	 // p  = x4 address
	// &p = the box address - pointer address
	// *p = x4 value


}

func setValue(str *string){
	*str = "Done"
}

func setValue2(str string) string{
	str = "Done-2"
	return str
}