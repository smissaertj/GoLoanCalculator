package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {

	// Parse provided parameters
	principal := flag.Float64("principal", 0, "The amount borrowed.")
	interest := flag.Float64("interest", 0, "The interest rate.")
	period := flag.Uint("period", 0, "Loan term; amount of time in months to pay off the loan.")
	amount := flag.Float64("amount", 0, "Monthly repayment amount.")
	flag.Parse()

	// Determine the missing parameter that needs to be calculated
	switch {
	case !isFlagPassed("principal"):
		// calculatePrincipal()
		fmt.Println("Calculating Principal")
	case !isFlagPassed("interest"):
		// calculateInterest()
		fmt.Println("Calculating Interest")
	case !isFlagPassed("period"):
		fmt.Println("Calculating Period")
		*period = calculatePeriod(*amount, *principal, *interest)
	case !isFlagPassed("amount"):
		// calculateAmount()
		fmt.Println("Calculating Amount")
	}

	fmt.Println("Principal", *principal, "Interest", *interest, "Period", *period, "Amount", *amount)
}

func isFlagPassed(flagName string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			found = true
		}
	})
	return found
}

func calculatePeriod(payment, principal, interestRate float64) uint {
	i := interestRate / (12 * 100) // Convert annual interest rate to monthly and to a decimal
	n := math.Log(payment/(payment-i*principal)) / math.Log(1+i)
	return uint(math.Ceil(n)) // Round up to the next whole number
}
