package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your Prices. Confirm each line with Enter. Type 'done' when finished: ")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "done" || price == "exit" {
			break
		}
		if price == "" {
			fmt.Println("Empty input. Please enter a valid price or type 'done' to finish.")
			continue
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
