package main

import (
	"awesomeProject2/ProGo/12/store"
	"awesomeProject2/ProGo/12/store/cart"
	. "awesomeProject2/ProGo/12/store/fmt"
	"fmt"
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

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
	fmt.Println("Price:", ToCurrency(product.Price()))

}
