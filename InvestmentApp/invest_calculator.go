package main

import (
	"fmt"
	"math"
)

// Instead of calling varible like this: var investmentAmount float64 = 1000 you can call it like this: investmentAmount := 1000 this is shorthand for declaring and initializing a variable.
// This is called short variable declaration.
// Defining varibles in same line is called multi-variable declaration.
// You can also define the type of the variable in the end like this: investmentAmount, expectedReturnRate, years := 1000, 5.5, 10

func main() {
	// Declare and initialize the investment amount
	var investmentAmount float64 = 1000

	// Declare and initialize the expected return rate
	var expectedReturnRate float64 = 5.5

	// Declare and initialize the number of years
	var years float64 = 10

	// Declare and initialize the inflation rate as a constant
	const inflationRate float64 = 3.5

	// Calculate the future value of the investment
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// Calculate the inflation-adjusted future value of the investment
	inflationAdjustedFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Print the future value of the investment
	fmt.Println("Future value of investment is ", futureValue)

	// Print the inflation-adjusted future value of the investment
	fmt.Println("Inflation adjusted future value of investment is ", inflationAdjustedFutureValue)
}
