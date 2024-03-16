package main

import "fmt"

func main() {
	revenueAmount := getUserInput("Revenue: ")
	expensesAmount := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningsBeforeTaxes, profit, ratio := calculateFinancials(revenueAmount, expensesAmount, taxRate)

	fmt.Println("EBT is ", earningsBeforeTaxes)
	fmt.Println("Your profit is ", profit)
	fmt.Println("Your ratio is ", ratio)
}

func calculateFinancials(revenueAmount, expensesAmount, taxRate float64) (float64, float64, float64) {
	ebt := revenueAmount - expensesAmount
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
