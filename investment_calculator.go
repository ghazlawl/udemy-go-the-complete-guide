package main

import (
	"fmt"
	"math"
)

func investment_calculator() {
	const inflationRate float64 = 2.5

	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var numYears float64 = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("# of Years: ")
	fmt.Scan(&numYears)

	var futureValue = investmentAmount * math.Pow(1+(expectedReturnRate/100), numYears)
	var futureValueAdjusted = futureValue / math.Pow(1+(inflationRate/100), numYears)

	fmt.Println(futureValue)
	fmt.Println(futureValueAdjusted)
}
