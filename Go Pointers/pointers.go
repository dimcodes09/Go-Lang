package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age
	fmt.Println("Age:", agePointer)

	 editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int){ //int the age stored here is teh copy which is created by go]
	// return *age - 18
	*age =  *age - 18
}
