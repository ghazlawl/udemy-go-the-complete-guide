package main

import "fmt"

func profileCalculator() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("EBT: $%.2f\n", ebt)
	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

/**
 * Prompts the user for input and returns the value as a float64.
 *
 * @param text The label to display when prompting the user.
 *
 * @return The user's input as a float64.
 */
func getUserInput(text string) float64 {
	var value float64

	fmt.Printf("%s: ", text)
	fmt.Scan(&value)

	return value
}

/**
 * Calculates the Earnings Before Tax (EBT), profit, and ratio based on revenue, expenses, and tax rate.
 *
 * @param revenue The total revenue.
 * @param expenses The total expenses.
 * @param taxRate The tax rate as a percentage.
 *
 * @return ebt The earnings before tax.
 * @return profit The profit after tax.
 * @return ratio The ratio of EBT to profit.
 */
func calculateProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - (taxRate / 100))
	ratio = ebt / profit

	return ebt, profit, ratio
}
