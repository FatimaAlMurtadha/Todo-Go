package main

import (
	"fmt"
	
)

type Points struct {
	x int32
	y int32	
	z bool
}
type Circle struct{
	radious float32
	center *Points
}

func StructsAndCustomTypes(){
	var x1 Points = Points{10 , 17, false}
	var x2  Points = Points{x: -9, y: 100, z: true}

	fmt.Println(x1, "  ",x2)
	fmt.Println(x1.x, x1.y, "  ",x2.x, x2.y)

	x1.x = 80
	x2.y = -20
	fmt.Println("\nAfter: \n",x1, "\n",x2)

	p1 := &Points{x: 13, y:40}
	fmt.Println(p1)
	fmt.Println(*p1)

	changeVal(p1)
	fmt.Println(p1)

	// ----------------------------------------------

	p2 := &Points{x: 5, y:8}
	c1 := Circle{5.12, p2}

	fmt.Println(c1)
	fmt.Println(c1.center)
	

}

// Send the value
func changeVal(pt *Points){
	pt.x += 50 
}