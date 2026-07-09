package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age
	fmt.Println("Age:", agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int{ //the age stored here is teh copy which is created by go]
	return *age - 18
}
