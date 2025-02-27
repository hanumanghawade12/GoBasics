package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// investmentAmount := 1000.0
	// annualInterestRate := 5.5
	var investmentAmount float64
	var annualInterestRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter annual interest rate in percentage: ")
	fmt.Scan(&annualInterestRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	// futureInvestmentValue := investmentAmount * math.Pow((1+annualInterestRate/100), years)
	// futurerealValue := futureInvestmentValue / math.Pow((1+inflationRate/100), years)
	futureInvestmentValue, futurerealValue := calculateFutureValues(investmentAmount, annualInterestRate, years)

	fmt.Printf("future value is $%.2f\n", futureInvestmentValue)
	fmt.Printf("future real value is $%.2f\n", futurerealValue)

}

func calculateFutureValues(investmentAmount float64, annualInterestRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+annualInterestRate/100), years)
	frv := fv / math.Pow((1+inflationRate/100), years)
	return fv, frv
}
