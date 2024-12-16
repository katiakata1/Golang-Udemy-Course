package main

import "fmt"

func main() {
	revenue := userInput("Revenue")
	expenses := userInput("Expenses")
	taxRate := userInput("Tax Rate")

	ebt, profit, ratio := calculation(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
}

func userInput(value string) (output float64) {
	fmt.Printf("%v: ", value)
	fmt.Scan(&output)
	return output
}

func calculation(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
