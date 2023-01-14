package main

import (
	"awesomeProject2/ProGo/12/store"
	"awesomeProject2/ProGo/12/store/cart"
	. "awesomeProject2/ProGo/12/store/fmt"
	"github.com/fatih/color"
)

func main() {

	// product := store.Product{
	// 	Name:     "Kayak",
	// 	Category: "Watersports",
	// 	price:    279,
	// }

	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	//
	//fmt.Println("Name:", cart.CustomerName)
	//fmt.Println("Total:", ToCurrency(cart.GetTotal()))
	//
	color.Green("Name: " + cart.CustomerName)
	color.Magenta("Total: " + ToCurrency(cart.GetTotal()))
}
