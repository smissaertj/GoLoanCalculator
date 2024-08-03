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
	payment := flag.Float64("payment", 0, "Monthly repayment amount.")
	periods := flag.Uint("periods", 0, "Loan term; amount of time in months to pay off the loan.")
	flag.Parse()

	// Determine the missing parameter that needs to be calculated
	switch {
	case !isFlagPassed("principal"):
		*principal = calculatePrincipal(*interest, *payment, *periods)
		fmt.Printf("Your loan principal is = %f!", *principal)
	case !isFlagPassed("interest"):
		fmt.Println("Please provide an annual interest rate!")
	case !isFlagPassed("periods"):
		*periods = calculatePeriods(*payment, *principal, *interest)
		fmt.Printf("Your loan period is = %d months!", *periods)
	case !isFlagPassed("payment"):
		*payment = calculatePayment(*principal, *interest, *periods)
		fmt.Printf("Your monthly payment = %f!", *payment)
	}

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
	return math.Round(numerator / denominator)
}

func calculatePeriods(principal, interest, payment float64) uint {
	i := convertInterest(interest)
	n := math.Log(payment/(payment-i*principal)) / math.Log(1+i)
	return uint(math.Ceil(n))
}

func calculatePayment(principal, interest float64, periods uint) float64 {
	i := convertInterest(interest)
	numerator := i * math.Pow(1+i, float64(periods))
	denominator := math.Pow(1+i, float64(periods)) - 1
	return math.Ceil(principal * (numerator / denominator))
}
