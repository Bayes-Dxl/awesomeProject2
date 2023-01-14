package main

import "fmt"

//type Expense interface {    getName() string    getCost(annual bool) float64}
func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
}
