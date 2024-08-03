package main

import (
	"flag"
	"fmt"
)

func main() {

	// Parse provided parameters
	principal := flag.Uint64("principal", 0, "The amount borrowed.")
	interest := flag.Uint("interest", 0, "The interest rate.")
	period := flag.Uint("period", 0, "Loan term; amount of time in months to pay off the loan.")
	amount := flag.Uint64("amount", 0, "Monthly repayment amount.")
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
		// calculatePeriod()
		fmt.Println("Calculating Period")
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
