package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {

	// Parse provided parameters
	paymentType := flag.String("type", "", "Type of payment: 'annuity' or 'diff'.")
	principal := flag.Float64("principal", 0, "The amount borrowed.")
	interest := flag.Float64("interest", 0, "The interest rate.")
	payment := flag.Float64("payment", 0, "Monthly repayment amount.")
	periods := flag.Uint("periods", 0, "Loan term; amount of time in months to pay off the loan.")
	flag.Parse()

	// Determine the missing parameter that needs to be calculated
	switch {
	case *paymentType == "" && *paymentType != "annuity" && *paymentType != "diff":
		hasInvalidParameters()
	case *paymentType == "annuity":
		// We can't calculate the interest, so it always needs to be provided
		if !isFlagPassed("principal") || !isFlagPassed("interest") || !isFlagPassed("periods") {
			hasInvalidParameters()
		}
	case *paymentType == "diff":
		// We can't calculate the principal or months, so a combination with the 'payment' flag is invalid
		if isFlagPassed("payment") || !isFlagPassed("principal") || !isFlagPassed("interest") || !isFlagPassed("periods") {
			hasInvalidParameters()
		}
	case !isFlagPassed("principal"):
		*principal = calculatePrincipal(*interest, *payment, *periods)
		fmt.Printf("Your loan principal is = %.0f!", *principal)
	case !isFlagPassed("interest"):
		fmt.Println("Please provide an annual interest rate!")
	case !isFlagPassed("periods"):
		*periods = calculatePeriods(*principal, *interest, *payment)
		formattedPeriod := formatMonthsToYearsAndMonths(int(*periods))
		fmt.Printf("It will take %s to repay this loan!", formattedPeriod)
	case !isFlagPassed("payment"):
		*payment = calculatePayment(*principal, *interest, *periods)
		fmt.Printf("Your monthly payment = %.0f!", *payment)
	}
}

func hasInvalidParameters() {
	fmt.Println("Incorrect parameters")
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

func formatMonthsToYearsAndMonths(totalMonths int) string {
	years := totalMonths / 12
	remainingMonths := totalMonths % 12

	if years == 0 {
		if remainingMonths == 1 {
			return "1 month"
		}
		return fmt.Sprintf("%d months", remainingMonths)
	}

	if years == 1 && remainingMonths == 0 {
		return "1 year"
	}

	yearString := "year"
	if years > 1 {
		yearString = "years"
	}

	if remainingMonths == 0 {
		return fmt.Sprintf("%d %s", years, yearString)
	}

	monthString := "month"
	if remainingMonths > 1 {
		monthString = "months"
	}

	return fmt.Sprintf("%d %s and %d %s", years, yearString, remainingMonths, monthString)
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
