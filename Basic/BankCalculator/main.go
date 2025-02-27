package main

import (
	"fmt"

	"exmaple.com/bank/fileops"
)

var filename = "bank_balance.txt"

func main() {
	bankBalance, err := fileops.GetFloatFromFile(filename)
	if err != nil {
		fmt.Println("Error getting balance from file")
	}
	fmt.Println("Welcome to the Bank Calculator")
	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Printf("Your bank balance is $%.2f\n", bankBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")
				return
			}
			bankBalance += depositAmount
			fmt.Printf("Your bank balance is $%.2f\n", bankBalance)
			fileops.WriteFloatToFile(bankBalance, filename)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount")
				return
			}
			if withdrawAmount > bankBalance {
				fmt.Println("Insufficient funds")
				return
			}
			bankBalance -= withdrawAmount
			fmt.Printf("Your bank balance is $%.2f\n", bankBalance)
			fileops.WriteFloatToFile(bankBalance, filename)

		} else if choice == 4 {
			fmt.Println("Exiting the Bank Calculator")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
