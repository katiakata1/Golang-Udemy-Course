package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer) // here we use a ponter to age value and passing it to getAdultYears which expect as pointer (*int)
	fmt.Println(age)                // getAdultYears mutated our original age value through agePointer, therefore now we can use age as output
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
	// *age is a direferencing of age in (age *int). Because we specified *int, age is automtically a pointer, therefore we need to make a real value to do a calculation
	// otherwise, as age will not be allowed to perform -18
	// this function also doesn't return anything therefore we need to remove int
}
