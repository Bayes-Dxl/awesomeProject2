package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax(0.2, 100))
}
func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}
func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	kayak.printDetails()
}
