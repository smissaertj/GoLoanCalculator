package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Parse provided parameters
	paymentType := flag.String("type", "", "Type of payment: 'annuity' or 'diff'.")
	principal := flag.Float64("principal", 0, "The amount borrowed.")
	interest := flag.Float64("interest", 0, "The interest rate.")
	payment := flag.Float64("payment", 0, "Monthly repayment amount.")
	periods := flag.Uint("periods", 0, "Loan term; amount of time in months to pay off the loan.")
	flag.Parse()

	switch {
	case *paymentType == "" && *paymentType != "annuity" && *paymentType != "diff":
		hasInvalidParameters()

	case *paymentType == "annuity":
		// Takes 3 out of 4 remaining parameters
		// interest -> required
		if !isPositiveFlagPassed("interest") {
			hasInvalidParameters()

		} else if !isPositiveFlagPassed("principal") && isPositiveFlagPassed("periods") && isPositiveFlagPassed("payment") {
			*principal = calculatePrincipal(*interest, *payment, *periods)
			overPayment := calculateOverpayment(*payment, *principal, *periods)
			fmt.Printf("Your loan principal is = %.0f!\nOverpayment = %0.f", *principal, overPayment)

		} else if !isPositiveFlagPassed("periods") && isPositiveFlagPassed("principal") && isPositiveFlagPassed("payment") {
			*periods = calculatePeriods(*principal, *interest, *payment)
			formattedPeriod := formatMonthsToYearsAndMonths(int(*periods))
			overPayment := calculateOverpayment(*payment, *principal, *periods)
			fmt.Printf("It will take %s to repay this loan!\nOverpayment = %.0f", formattedPeriod, overPayment)

		} else if !isPositiveFlagPassed("payment") && isPositiveFlagPassed("principal") && isPositiveFlagPassed("periods") {
			annuityPayment := calculateAnnuityPayment()
			fmt.Printf("Your annuity payment = %.2f\n", annuityPayment)
			fmt.Printf("Overpayment = %.2f", annuityPayment)

		} else {
			hasInvalidParameters()
		}

	case *paymentType == "diff":
		// Takes 3 out of 4 remaining parameters
		// payment -> invalid flag
		if isPositiveFlagPassed("payment") || (!isPositiveFlagPassed("principal") || !isPositiveFlagPassed("interest") || !isPositiveFlagPassed("periods")) {
			hasInvalidParameters()

		} else {
			*payment = calculatePayment(*principal, *interest, *periods)
			fmt.Printf("Your monthly payment = %.0f!", *payment)
		}

	}
}

func hasInvalidParameters() {
	fmt.Println("Incorrect parameters")
}

func isPositiveFlagValue(f *flag.Flag) bool {
	value, _ := strconv.ParseFloat(f.Value.String(), 64)
	return value > 0
}

func isPositiveFlagPassed(flagName string) bool {
	isFoundAndPositive := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			isFoundAndPositive = isPositiveFlagValue(f)
		}
	})
	return isFoundAndPositive
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

func calculateOverpayment(payment, principal float64, periods uint) float64 {
	return (payment * float64(periods)) - principal
}

func calculatePrincipal(interest, amount float64, period uint) float64 {
	i := convertInterest(interest)
	numerator := amount
	denominator := (i * math.Pow(1+i, float64(period))) / (math.Pow(1+i, float64(period)) - 1)
	return math.Floor(numerator / denominator)
}

func calculatePeriods(principal, interest, payment float64) uint {
	i := convertInterest(interest)
	n := math.Log(payment/(payment-i*principal)) / math.Log(1+i)
	return uint(math.Ceil(n))
}

func calculateAnnuityPayment() float64 {
	return 0.0
}

func calculatePayment(principal, interest float64, periods uint) float64 {
	i := convertInterest(interest)
	numerator := i * math.Pow(1+i, float64(periods))
	denominator := math.Pow(1+i, float64(periods)) - 1
	return math.Ceil(principal * (numerator / denominator))
}
