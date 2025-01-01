package main

import "fmt"

func main() {
	var date string
	var description string
	var amount float64

	date = "2024-12-18"
	description = "Groceries"
	amount = 123.45

	fmt.Printf("Date: %s\nDescription: %s\nAmount: $%.2f\n", date, description, amount)
}

