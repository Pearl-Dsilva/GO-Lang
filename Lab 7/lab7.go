package main

import (
	"errors"
	"fmt"
)

// This structure represents a product in the supermarket.
type Product struct {
	ID    int
	Name  string
	Price float64
}

// CartItem is a structure that represents an item in the shopping cart.
type CartItem struct {
	Product  Product
	Quantity int
}

// ShoppingCart represents a customer's shopping cart...
type ShoppingCart struct {
	Items []CartItem
}

/*
Interfaces:
Customer interface defines methods related to a supermarket customer.
*/
type Customer interface {
	AddToCart(productID int, quantity int, productCatalog map[int]Product) error
	Checkout() float64
}

// Cashier interface defines methods related to a supermarket cashier.
type Cashier interface {
	ScanItem(item CartItem)
	CalculateTotal() float64
	ReceivePayment(amount float64) error
}

// implementation of RegularCustomer structure
type RegularCustomer struct {
	ShoppingCart
}

func (c *RegularCustomer) AddToCart(productID int, quantity int, productCatalog map[int]Product) error {
	product, exists := productCatalog[productID]
	if !exists {
		return errors.New("product not found")
	}

	item := &CartItem{Product: product, Quantity: quantity}
	c.Items = append(c.Items, *item) // Dereference the pointer before appending
	fmt.Printf("Added %d x %s to the cart.\n", quantity, product.Name)
	return nil
}

func (c *RegularCustomer) Checkout() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

// SupermarketCashier implementation
type SupermarketCashier struct {
	ScannedItems []CartItem
	Total        float64
}

func (s *SupermarketCashier) ScanItem(item *CartItem) {
	s.ScannedItems = append(s.ScannedItems, *item) // Dereference the pointer before appending
	s.Total += item.Product.Price * float64(item.Quantity)
	fmt.Printf("Scanned %d x %s.\n", item.Quantity, item.Product.Name)
}

func (s *SupermarketCashier) CalculateTotal() float64 {
	return s.Total
}

func (s *SupermarketCashier) ReceivePayment(amount float64) error {
	if amount < s.Total {
		return errors.New("insufficient payment")
	}
	return nil
}

func main() {
	// Initialize product catalog
	productCatalog := map[int]Product{
		1: {ID: 1, Name: "Bread", Price: 50.0},
		2: {ID: 2, Name: "Milk", Price: 30.5},
		3: {ID: 3, Name: "Tomatoes", Price: 100.0},
		4: {ID: 4, Name: "Potatos", Price: 70.0},
		5: {ID: 5, Name: "Onions", Price: 60.0},
	}

	customer := &RegularCustomer{}
	cashier := &SupermarketCashier{}

	for {
		fmt.Println("\n1. Add item to cart")
		fmt.Println("2. View cart and checkout")
		fmt.Println("3. View the product catalog")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// case to add items to cart
			fmt.Print("Enter product ID: ")
			var productID int
			_, err := fmt.Scan(&productID)
			if err != nil {
				fmt.Println("Invalid product ID. Please enter a valid number.")
				continue
			}

			fmt.Print("Enter quantity: ")
			var quantity int
			_, err = fmt.Scan(&quantity)
			if err != nil {
				fmt.Println("Invalid quantity. Please enter a valid number.")
				continue
			}

			if err := customer.AddToCart(productID, quantity, productCatalog); err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			// case to handle view cart and checkout
			total := customer.Checkout()
			fmt.Printf("Cart Total: Rs. %.2f\n", total)

			for _, item := range customer.Items {
				cashier.ScanItem(&item) // Pass the address of the CartItem
			}

			cashierTotal := cashier.CalculateTotal()
			fmt.Printf("Cashier's Total: Rs. %.2f\n", cashierTotal)

			fmt.Print("Enter payment amount: ")
			var paymentAmount float64
			_, err := fmt.Scan(&paymentAmount)
			if err != nil {
				fmt.Println("Invalid payment amount. Please enter a valid number.")
				continue
			}

			if err := cashier.ReceivePayment(paymentAmount); err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Please provide sufficient funds.")
			} else {
				difference := paymentAmount - cashierTotal
				fmt.Println("Payment successful. Thank you!")
				if difference > 0.0 {
					fmt.Println("Collect your balance: Rs. ", difference)
				}
			}

			// reset the customer's cart
			customer.Items = nil

		case 3:
			fmt.Println("Product Catalog:")
			for id, product := range productCatalog {
				fmt.Printf("%d. %s - $%.2f\n", id, product.Name, product.Price)
			}

		case 4:
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
