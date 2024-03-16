package main

import (
	"fmt"
	"math"
)

func main() {
	const INFLATION_RATE float64 = 3.2
	var investmentAmount float64 = 10000.00
	var expectedReturn float64 = 5.5
	var years float64 = 10

	//////  GET VALUES   //////////
	fmt.Print("Enter the amount you'd like to invest: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Enter the amount of years: ")
	fmt.Scan(&years)

	//////  CALCULATIONS   //////////

	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	//////  PRINT RESULTS   //////////

	fmt.Println("Investment Amount: $", investmentAmount)
	fmt.Println("Expected Return: ", expectedReturn, "%")
	fmt.Println("Years: ", years)
	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Real Value: %.2f\n", futureRealValue)
}
