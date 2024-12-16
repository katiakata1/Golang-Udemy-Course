package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	// we are making go aware that we want to return error, which is a special type in go
	data, err := os.ReadFile(accountBalanceFile)

	// err not equals to nill means that there is an error and we are checking this now
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	// now we need to return two values as we specified here (float64, error). We return nil?
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// values by order: file name, bytes???, and permissions on the file
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	// our function now returns two values so we need to adjust for it
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't continue, sorry")
	}
	fmt.Println("Welcome to go bank")

	// for i := 0; i < 200; i++
	// to loop forever
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("How much would you like to withdraw?: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			accountBalance -= withdrawAmount

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than you have")
				continue
			}

			fmt.Println("New balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing our bank")
			return
			// break
		}
	}
}
