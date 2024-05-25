package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Product is a struct that represents a product in the supermarket.
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// MarshalJSON customizes JSON marshaling for the Product struct.
func (p *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	})
}

// UnmarshalJSON customizes JSON unmarshaling for the Product struct.
func (p *Product) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	p.Name = temp.Name
	p.Price = temp.Price
	p.Stock = temp.Stock
	return nil
}

func main() {
	var choice int
	for choice != 3 {
		fmt.Println("\n1. Marshal JSON")
		fmt.Println("2. Unmarshal JSON")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			var price float64
			var stock int

			fmt.Print("Enter product name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter product price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter product stock: ")
			fmt.Scanln(&stock)

			product := Product{Name: name, Price: price, Stock: stock}

			// Reads the product detaills, marshals it to JSON and stores it in a file.
			jsonBytes, err := json.Marshal(&product)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return
			}
			if err := os.WriteFile("product.json", jsonBytes, 0644); err != nil {
				fmt.Println("Error writing JSON to file:", err)
				return
			}
			fmt.Println("Marshalled JSON stored in product.json")

		case 2:
			// Reads JSON from file and unmarshals it into a product.
			jsonBytes, err := os.ReadFile("product.json")
			if err != nil {
				fmt.Println("Error reading JSON from file:", err)
				return
			}
			var product Product
			if err := json.Unmarshal(jsonBytes, &product); err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			fmt.Println("Unmarshalled Product:", product)

		case 3:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid choice")
		}
	}
}
