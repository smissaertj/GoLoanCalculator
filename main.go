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
	payment := flag.Float64("amount", 0, "Monthly repayment amount.")
	period := flag.Uint("period", 0, "Loan term; amount of time in months to pay off the loan.")
	flag.Parse()

	// Determine the missing parameter that needs to be calculated
	switch {
	case !isFlagPassed("principal"):
		*principal = math.Round(calculatePrincipal(*interest, *payment, *period))
		fmt.Printf("Your loan principal is = %f!", *principal)
	case !isFlagPassed("interest"):
		// calculateInterest()
		fmt.Println("Calculating Interest")
	case !isFlagPassed("period"):
		*period = calculatePeriod(*payment, *principal, *interest)
		fmt.Printf("Your loan period is = %d months!", *period)
	case !isFlagPassed("amount"):
		// calculateAmount()
		fmt.Println("Calculating Amount")
	}

	fmt.Println("Principal", *principal, "Interest", *interest, "Period", *period, "Payment", *payment)
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

func convertInterest(interest float64) float64 {
	return interest / (12 * 100) // Convert annual interest rate to monthly and to a decimal
}

func calculatePrincipal(interest, amount float64, period uint) float64 {
	i := convertInterest(interest)
	numerator := amount
	denominator := (i * math.Pow(1+i, float64(period))) / (math.Pow(1+i, float64(period)) - 1)
	return numerator / denominator
}

func calculatePeriod(principal, interest, payment float64) uint {
	i := convertInterest(interest)
	n := math.Log(payment/(payment-i*principal)) / math.Log(1+i)
	return uint(math.Ceil(n)) // Round up to the next whole number
}
