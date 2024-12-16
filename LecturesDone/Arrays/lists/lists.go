package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices = append(prices, 5.99, 12.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	// ... in go means that it can go inside that list, take every itim and separate it by comma
	// if you used (prices, discountPrices) only, you would get an error because prices is a list of items and not list of lists
	// therefore, you cannot add list of items to items.
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	// here we tell that we create a productNames variable contains 4 values which is an array of strings
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 3.99, 46.99, 56.90}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[0])

// 	// when you create a slice, the copy of slice is not created
// 	// it is just a reference to the original array
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	// here you will see only 3.99 because :1 means that it come to second element (index1) and exclude it
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	// you can also identify length and capacity of the array using two functions below
// 	// len is number of items in a slice or array
// 	// cap is number of items in a previous slice/array
// 	// that is why cap(highlightedPrices) = 3 because featuredPrices was cut from left, and the rest 3 on the rights are considered to be a capacity
// 	// that is why cap(featuredPrices) is 3, because prices was cut from the left
// 	// you can select more items to the right, but never to the left
// 	fmt.Println(len(prices), cap(prices))
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	// you can reslice the last
// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
