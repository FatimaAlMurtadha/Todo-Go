package main

import (
	"fmt"
)

type Students struct{
	name string
	age int
	grades []int

}

func MyMethods(){
	// Struct Methods 

	s1 := Students{"Fatima", 33, []int {100, 90, 95}}

	fmt.Println(s1.getAge())
	fmt.Println("The grades are: ",s1.getGrades())

	fmt.Println(s1) // {Fatima 33 [100 90 95]}

	s1.setAge(70) 
	fmt.Println(s1) // {Fatima 33 [100 90 95]}
    // Didn't change - because we didn't send the value as a pointer

	s1.setAge2(70)
	fmt.Println(s1) // {Fatima 70 [100 90 95]}

	fmt.Println(s1.getAverage())

	s2 := Students{"Safa", 15, []int{50, 40, 66}}

	fmt.Println(s2.getAverage())

}

func (s Students) getAge() string{

	return fmt.Sprint("The age is: ", s.age)
}

func (s Students) setAge(age int){
	s.age = age
}

func (s *Students) setAge2(age int){
	s.age = age
}

func (s Students) getGrades() []int{

	return s.grades
}

func (s Students) getAverage() float32{
	sum := 0
	for _, i := range s.grades{
		sum += i
	}
	return float32(sum) / float32(len(s.grades))
}

