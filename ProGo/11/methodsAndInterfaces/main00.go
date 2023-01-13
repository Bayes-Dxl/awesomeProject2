package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

//func calcTotal(expenses []Expense) (total float64) {
//	for _, item := range expenses {
//		total += item.getCost(true)
//	}
//	return
//}

//type EXPENSE []Expense
//
//func (p EXPENSE) calcTotal1() (total float64) {
//	for i, item := range p {
//		total += item.getCost(true)
//		fmt.Println(i)
//	}
//	return
//}
//
//type Account struct {
//	accountNumber int
//	//expenses      []Expense
//	expenses EXPENSE
//}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}

func main00() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}

	for _, item := range data {
		processItem(item)
	}
}

/*
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}
	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:",
		insurance.monthlyFee*float64(insurance.durationMonths))

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total: ", calcTotal(expenses))
	fmt.Println("------------------------------------")

	account := Account{
		accountNumber: 12345,
		expenses: EXPENSE{ //  []Expense
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))
	fmt.Println("Total:", account.expenses.calcTotal1())

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))

	product1 := Product{"Kayak", "Watersports", 275}
	var expense1 Expense = &product1
	product1.price = 100
	fmt.Println("Product field value:", product1.price)
	fmt.Println("Expense method result:", expense1.getCost(false))
*/

//expenses := []Expense{
//	Service{"Boat Cover", 12, 89.50, []string{}},
//	Service{"Paddle Protect", 12, 8, []string{}},
//}
//for _, expense := range expenses {
//	s := expense.(Service)
//	fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
//}

//expenses := []Expense{
//	Service{"Boat Cover", 12, 89.50, []string{}},
//	Service{"Paddle Protect", 12, 8, []string{}},
//	&Product{"Kayak", "Watersports", 275},
//}
//for _, expense := range expenses1 {
//	s, ok := expense.(Service)
//	if ok {
//		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
//	} else {
//		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
//
//	}
//
//}

//for _, expense := range expenses {
//	switch value := expense.(type) {
//	case Service:
//		fmt.Println("Service:", value.description, "Price:",
//			value.monthlyFee*float64(value.durationMonths))
//	case *Product:
//		fmt.Println("Product:", value.name, "Price:", value.price)
//	default:
//		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
//	}
//}
