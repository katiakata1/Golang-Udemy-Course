package main

import (
	"errors"
	"fmt"
	"os"
)

func writeBalanceToFile(ebt, profit, ratio float64) {
	// Sprintf allows some formatting rather than Sprint
	results := fmt.Sprintf("EBT: %.1f\nPROFIT: %.1f\nRATIO: %.1f\n", ebt, profit, ratio)
	os.WriteFile("balance.txt", []byte(results), 0644)
}

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeBalanceToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("the value should not be less or equal to zero")
	}
	return userInput, nil
}
