// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	const inflationRate = 2.5
// 	var investmentAmount float64
// 	var years float64
// 	expectedReturnRate := 5.5

// 	fmt.Print("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Print("Years: ")
// 	fmt.Scan(&years)

// 	fmt.Print("Expected return: ")
// 	fmt.Scan(&expectedReturnRate)

// 	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

// 	formattteFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
// 	formattteRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)

// 	// fmt.Println(futureValue)
// 	// fmt.Println("Futute value: ", futureValue)
// 	// fmt.Printf("Future value: %.0f, %f", futureValue, futureRealValue)
// 	// fmt.Println(futureRealValue)
// 	fmt.Print(formattteFV, formattteRFV)
// }
