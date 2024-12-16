package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 67)

	// here numbers... takes this slice []int{1, 10, 15} and turn every value into a list like in sumup(1, 10, 15, 67)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// THIS IS VARIATIC FUNCTION
// sumup can be called with a list of separate values
// it can have as many parameter values as you want
// as long as every single values is of type int
// startingValue int will be 1 in (1, 10, 15, 67); the rest will be numbers ...int
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
