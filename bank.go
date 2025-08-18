package main

import (
	"fmt"

	"example.com/bank/io"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFilename = "balance.txt"

func main() {
	var accountBalance, err = io.ReadFloatFromFile(accountBalanceFilename)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Printf("Your initial balance is: $%.2f\n", accountBalance)
	fmt.Printf("Reach us 24/7 at %s\n", randomdata.Email())

	for {
		displayOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is: $%.2f\n", accountBalance)
		}

		if choice == 2 {
			var depositAmount float64

			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero!")
				continue
			}

			accountBalance += depositAmount

			fmt.Printf("You deposited: $%.2f\n", depositAmount)
			fmt.Printf("Your new balance is: $%.2f\n", accountBalance)

			io.WriteFloatToFile(accountBalance, accountBalanceFilename)
		}

		if choice == 3 {
			var withdrawAmount float64

			fmt.Print("Enter amount to withdraw: ")
			fmt.Scanln(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than zero!")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				accountBalance -= withdrawAmount

				fmt.Printf("You withdrew: $%.2f\n", withdrawAmount)
				fmt.Printf("Your new balance is: $%.2f\n", accountBalance)

				io.WriteFloatToFile(accountBalance, accountBalanceFilename)
			}
		}

		if choice == 4 {
			fmt.Println("Thank you for using Go Bank! Goodbye!")
			break
		}
	}
}
