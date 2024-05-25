package main

/*
Domain: Supermarket Management
Scenerio:
System for the supemarket to manage it's stocks/ inventory
*/

import (
	"errors"
	"fmt"
)

// structure is created with differnt data types like int, string, float and uint
type Inventory struct {
	id       int
	item     string
	price    float32
	quantity uint
}

// size is used to give a max size for array of structures. In order to simplify the program max size is declared as 10
const size int = 10

// variable for front and rear to traverse through Queue.
var front int = -1
var rear int = -1
var stock [size]Inventory

func main() {
	var choice int = 0
	for choice != 6 {
		fmt.Println("Enter your choice")
		fmt.Println("1. Add items to the inventory")
		fmt.Println("2. Get total cost of all the items in the inventory")
		fmt.Println("3. Display inventory")
		fmt.Println("4. Display top 3 items from inventory")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addItem()
		case 2:
			getTotalCost()
		case 3:
			displayInventory(front, rear)
		case 4:
			topN(stock[0:3])
		case 6:
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}

}

/*
Adds the items into the inventory
*/
func addItem() {
	if rear+1 == size {
		fmt.Println("Can't add more items")
		return
	} else if rear == -1 && front == -1 {
		rear = 0
		front = 0
	} else {
		rear++
	}
	var (
		id       int
		item     string
		price    float32
		quantity uint
	)
	fmt.Println("Enter the id of the item to add: ")
	fmt.Scan(&id)
	fmt.Println("Enter the name of the item to add: ")
	fmt.Scan(&item)
	fmt.Println("Enter the price of the item ", item, ": ")
	fmt.Scan(&price)
	fmt.Println("Enter the quantity of the item ", item, ": ")
	fmt.Scan(&quantity)

	stock[rear].id = id
	stock[rear].item = item
	stock[rear].price = price
	stock[rear].quantity = quantity
	fmt.Println("Item Added Successfully!")
}

/*
This function gets the total cost of the entire inventory
*/
func getTotalCost() {
	var total float32
	for _, o := range stock {
		total += (o.price * float32(o.quantity))
	}
	fmt.Println("The total cost of the inventory: ", total)
}

/*
Displays the inventory - items, price, quantity details of all the items in the inventory
*/
func displayInventory(front int, rear int) {
	if rear == -1 {
		fmt.Println("Inventory is empty")
		return
	} else {
		for i := front; i <= rear; i++ {
			fmt.Println("----------")
			fmt.Println("Id: ", stock[i].id)
			fmt.Println("Name: ", stock[i].item)
			fmt.Println("Price: ", stock[i].price)
			fmt.Println("Quantity: ", stock[i].quantity)
		}
	}
}

/*
Uses the concept od slices to show top 3 item details in inventory
*/
func topN(stockData []Inventory) {
	for _, v := range stockData {
		fmt.Println("----------")
		fmt.Println("Id: ", v.id)
		fmt.Println("Name: ", v.item)
		fmt.Println("Price: ", v.price)
		fmt.Println("Quantity: ", v.quantity)
	}
}

func lessThanShowError(stockData []Inventory, max uint) (res string, err error) {
	if stockData[0].quantity >= max {
		return "", errors.New("More than max")
	} else {
		return res, err
	}
}
