package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected return: ")
	outputText("Expected return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattteFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
	formattteRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)

	fmt.Print(formattteFV, formattteRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// Return means that the value which is generated with the function should be returned as a result
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
