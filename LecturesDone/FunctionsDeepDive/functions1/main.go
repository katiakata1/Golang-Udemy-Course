package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := tranformNumbers(&numbers, double)
	tripled := tranformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	tranformedNumbers := tranformNumbers(&numbers, transformFn1)
	moreTranformedNumbers := tranformNumbers(&moreNumbers, transformFn2)

	fmt.Println(tranformedNumbers)
	fmt.Println(moreTranformedNumbers)

}

func tranformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// you return a function here!
// transformFn is a type which is returned. We defined it the beginning
// to make this function more useful we want it to accept pointer to a slice and name it numbers
func getTransformerFunction(numbers *[]int) transformFn {
	// here we check if the first number in there equals to 1
	// only structs have this shortcut where you don't need to dereference them
	// Therefore in our case we need to dereference a pointer to get a value
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
	// if you did double(), you would execute double function and return a result
	// but with just double, you return a whole function

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
