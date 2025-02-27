package main

import (
	"errors"
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter revenue: ")
	// outputText("Enter revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Enter expenses: ")
	// outputText("Enter expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("Enter tax rate: ")
	// outputText("Enter tax rate: ")
	// fmt.Scan(&taxRate)
	// expensesBeforeTex := revenue - expenses
	// profit := expensesBeforeTex * (1 - (taxRate / 100))
	// ratio := profit / revenue * 100
	revenue, err1 := outputText("Enter revenue: ")
	expenses, err2 := outputText("Enter expenses: ")
	taxRate, err3 := outputText("Enter tax rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}
	expensesBeforeTex, profit, ratio := caclulateProfit(expenses, revenue, taxRate)

	fmt.Println("Profit before tax: ", expensesBeforeTex)
	fmt.Println("Profit after tax: ", profit)
	fmt.Println("Profit ratio: ", ratio)
}

func outputText(text string) (float64, error) {
	var inputValue float64
	fmt.Print(text)
	fmt.Scan(&inputValue)
	if inputValue <= 0 {
		return 0, errors.New("Invalid input")
	}

	return inputValue, nil
}

func caclulateProfit(expenses, revenue, taxRate float64) (float64, float64, float64) {
	expensesBeforeTex := revenue - expenses
	profit := expensesBeforeTex * (1 - (taxRate / 100))
	ratio := profit / revenue * 100
	return expensesBeforeTex, profit, ratio

}
