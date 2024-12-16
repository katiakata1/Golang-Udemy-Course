package main

import "fmt"

func main() {
	fact := factorial(3)
	fmt.Println(fact)
}

// Recursive Function
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// factorial of 5 -> 5 * 4 * 3 * 2 * 1 = 120
