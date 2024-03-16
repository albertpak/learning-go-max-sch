package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenueAmount, revenueAmountErr := getUserInput("Revenue: ")
	expensesAmount, expensesAmountErr := getUserInput("Expenses: ")
	taxRate, taxRateError := getUserInput("Tax Rate: ")

	if revenueAmountErr != nil || expensesAmountErr != nil || taxRateError != nil {
		fmt.Println("Must be positive number")
		return
	}

	earningsBeforeTaxes, profit, ratio := calculateFinancials(revenueAmount, expensesAmount, taxRate)

	fmt.Println("EBT is ", earningsBeforeTaxes)
	fmt.Println("Your profit is ", profit)
	fmt.Println("Your ratio is ", ratio)

	storeResults(earningsBeforeTaxes, profit, ratio)
}

func storeResults(earningsBeforeTaxes, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", earningsBeforeTaxes, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenueAmount, expensesAmount, taxRate float64) (float64, float64, float64) {
	ebt := revenueAmount - expensesAmount
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive")
	}

	return userInput, nil
}
