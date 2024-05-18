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
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println("Future value of investment is ", futureValue)
}
