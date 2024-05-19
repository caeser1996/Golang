package main

import "fmt"

func main() {
	var revenue float64 // Variable to store the revenue
	var expenses float64 // Variable to store the expenses
	var taxRate float64 // Variable to store the tax rate

	fmt.Print("Enter the revenue: ") // Prompting the user to enter the revenue
	fmt.Scan(&revenue) // Reading the revenue input from the user

	fmt.Print("Enter the expenses: ") // Prompting the user to enter the expenses
	fmt.Scan(&expenses) // Reading the expenses input from the user

	fmt.Print("Enter the tax rate: ") // Prompting the user to enter the tax rate
	fmt.Scan(&taxRate) // Reading the tax rate input from the user

	ebt := revenue - expenses // Calculating the earnings before tax
	profit := ebt * (1 - taxRate/100) // Calculating the profit after tax
	ratio := ebt / profit // Calculating the ratio of earnings before tax to profit

	fmt.Println("Earnings before tax is ", ebt) // Printing the earnings before tax
	fmt.Println("Profit is ", profit) // Printing the profit
	fmt.Println("Ratio is ", ratio) // Printing the ratio

	// inline print using printf and add everythign in one line
	fmt.Printf("Earnings before tax is %.2f, Profit is %.2f, Ratio is %.2f\n", ebt, profit, ratio)

}
