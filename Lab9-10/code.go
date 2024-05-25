package main

import (
	"fmt"
	"sync"
)

// Item structure represents a supermarket item
type Item struct {
	Name     string
	Quantity int
}

// Customer structure represents a customer's shopping cart
type Customer struct {
	ID    int
	Items []Item
	Total float64
}

// Transaction struct represents a completed transaction
type Transaction struct {
	CustomerID int
	Total      float64
}

/*
The checkout function, takes the customer structure containing id and list of items.
takes in a waitgroup and channel transactions that can only send values of type struct, transactions.
The function iterates through the list of items and multiplies with the corresponding item prices and sums it to make a total.
Sends the structure Transaction with values, customer id and total from previous step.
Notify the waitgroup that the current trasaction is completed/ done.
*/
func checkout(customer Customer, wg *sync.WaitGroup, transactions chan<- Transaction) {
	// Inventory
	itemPrices := map[string]float64{
		"Apples":  200,
		"Bananas": 75,
		"Milk":    30,
	}
	var total float64
	for _, item := range customer.Items {
		price := itemPrices[item.Name]
		total += float64(item.Quantity) * price
	}

	// Send the transaction details to the transactions channel
	transactions <- Transaction{customer.ID, total}

	// Notify wait group that this goroutine is done
	wg.Done()
}

func main() {
	// Consider multiple supermarket customers waiting for checkout
	customers := []Customer{
		{ID: 1, Items: []Item{{Name: "Apples", Quantity: 3}, {Name: "Bananas", Quantity: 2}}},
		{ID: 2, Items: []Item{{Name: "Milk", Quantity: 1}}},
		{ID: 3, Items: []Item{{Name: "Apples", Quantity: 5}, {Name: "Milk", Quantity: 2}}},
		{ID: 4, Items: []Item{{Name: "Bananas", Quantity: 5}, {Name: "Milk", Quantity: 2}}},
	}

	// Creating a channel to collect transactions using make
	transactions := make(chan Transaction)

	// Creating a WaitGroup to wait for all checkout goroutines to finish
	var wg sync.WaitGroup

	// Adding the number of customers to the WaitGroup
	wg.Add(len(customers))

	// Process each customer's checkout concurrently
	for _, customer := range customers {
		// Launching the goroutine to handle each customer's checkout
		go checkout(customer, &wg, transactions)
	}

	// Closing transactions channel when all checkouts are done
	go func() {
		wg.Wait()
		close(transactions)
	}()

	// collecting and printing transactions
	for transaction := range transactions {
		fmt.Printf("CustomerID: %d, Total: Rs. %.2f\n", transaction.CustomerID, transaction.Total)
	}
}
