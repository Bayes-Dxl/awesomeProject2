package main

import "fmt"

func func1() {
	if arg := recover(); arg != nil {
		if err, ok := arg.(error); ok {
			fmt.Println("Error:", err.Error())
		} else if str, ok := arg.(string); ok {
			fmt.Println("Message:", str)
		} else {
			fmt.Println("Panic recovered")
		}
	}
}

type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage{TerminalError: arg}
		}
		close(outChan)
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func main() {

	//recoveryFunc := func() {
	//	if arg := recover(); arg != nil {
	//		if err, ok := arg.(error); ok {
	//			fmt.Println("Error:", err.Error())
	//		} else if str, ok := arg.(string); ok {
	//			fmt.Println("Message:", str)
	//		} else {
	//			fmt.Println("Panic recovered")
	//		}
	//	}
	//}
	//defer recoveryFunc()
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}
}
