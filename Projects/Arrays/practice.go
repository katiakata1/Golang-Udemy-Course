package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := []string{"LD", "Gym", "crafting"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3)
	hobbiesFirst := hobbies[:2]
	// hobbiesSecond := hobbies[0:2]
	fmt.Println(hobbiesFirst)

	// 4)
	fmt.Println(cap(hobbiesFirst))
	// doesn't work like this
	// hobbiesFirst = hobbiesFirst[1:]
	// fmt.Println(cap(hobbiesFirst))

	// If you want to reslice the previous array you need to specify start and end
	hobbiesFirst = hobbiesFirst[1:3]
	fmt.Println(hobbiesFirst)

	// 5)
	courseGoals := []string{"Learn go", "don't get fired"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "I am the best"
	fmt.Println(courseGoals)

	// 7)
	courseGoals = append(courseGoals, "My third goal")
	fmt.Println(courseGoals)

	// 8)
	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-prodcut",
			"My first product",
			4.99,
		},
	}

	fmt.Println(products)

	// 9)
	newProduct := Product{
		"third-product",
		"A third product",
		15.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
