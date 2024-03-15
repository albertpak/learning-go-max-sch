package main

import "fmt"

func main() {
	var (
		revenueAmount  float64
		expensesAmount float64
		taxRate        float64
	)

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenueAmount)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expensesAmount)

	fmt.Print("What is your tax rate? ")
	fmt.Scan(&taxRate)

	earningsBeforeTaxes := revenueAmount - expensesAmount
	profit := earningsBeforeTaxes * (1 - taxRate/100)

	fmt.Println("EBT is ", earningsBeforeTaxes)
	fmt.Println("Your profit is ", profit)
}
