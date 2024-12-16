package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

// this function is called generic function add[]
// it accepts any type of value only if it is int, float64, string
// if you choose any type, some types don't allow summation, for example Structs, therefore go complains
// generic function allows to accept certain types and it gives back certain types
// therefore result in main() knows that these types will be returned and not any
func add[T int | float64 | string](a, b T) T {
	return a + b
}
