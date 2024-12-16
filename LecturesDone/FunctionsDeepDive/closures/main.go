package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// here you create a function which accepts int and returns int
	double := createTranformer(2)
	triple := createTranformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	// now you can use that double function you created and use it as input to transformNumbers
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// should produce function that takes int and returns int
// createTranformer should take parameters too
func createTranformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
