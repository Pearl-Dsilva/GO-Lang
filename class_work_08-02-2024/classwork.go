package main

import (
	"fmt"
)

func main() {
	var choice int
	for choice != 8 {
		fmt.Println("\n1. Q1")
		fmt.Println("2. Q2")
		fmt.Println("3. Q3")
		fmt.Println("4. Q4")
		fmt.Println("5. Q5")
		fmt.Println("6. Q6")
		fmt.Println("7. Q6.2")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			q1()
		case 2:
			q2()
		case 3:
			q3()
		case 4:
			q4()
		case 5:
			q5()
		case 6:
			q6()
		case 7:
			q6_2()
		case 8:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}

/*
Q1
You're developing an online store application in GoLang. As part of the application,
you need to keep track of various product details such as name, price, and quantity in stock.
Design a set of variables and assign values to represent a specific product in the inventory.
Ensure you use appropriate data types for each variable to accurately capture the information.
*/
func q1() {
	var name string
	var price float32
	var quantity int8

	fmt.Println("Enter the bleow details of the product\nName: ")
	fmt.Scan(&name)
	fmt.Println("Price: ")
	fmt.Scan(&price)
	fmt.Println("Quantity: ")
	fmt.Scan(&quantity)
	fmt.Println("Name: ", name, " Price: ", price, " Quantity: ", quantity)
}

/*
Q2
You're tasked with building a student information system in GoLang for a school.
Each student record needs to store details such as student ID, name, age, and grade.
Define variables to store the information of a single student and assign values accordingly.
Pay attention to selecting appropriate data types to represent each piece of information.
*/
func q2() {
	var id int8
	var name string
	var age int8
	var grade string

	fmt.Println("Enter the bleow details of the students\nID: ")
	fmt.Scan(&id)
	fmt.Println("Name: ")
	fmt.Scan(&name)
	fmt.Println("Age: ")
	fmt.Scan(&age)
	fmt.Println("Grade: ")
	fmt.Scan(&grade)
	fmt.Println("ID: ", id, "Name: ", name, " Age: ", age, " Grade: ", grade)
}

/*
Q3
You're developing a weather monitoring system in GoLang for a research institute.
The system needs to store data about temperature, humidity, and wind speed.
Define variables to hold these measurements for a specific location at a given point in time.
Ensure you choose suitable data types for representing numerical measurements accurately.
*/
func q3() {
	var temparature float32
	var humidity float32
	var windSpeed float32

	fmt.Println("Enter the bleow details\nTemparature: ")
	fmt.Scan(&temparature)
	fmt.Println("himidity: ")
	fmt.Scan(&humidity)
	fmt.Println("wind speed: ")
	fmt.Scan(&windSpeed)
	fmt.Println("Temparature: ", temparature, "Humidity: ", humidity, " Wind Speed: ", windSpeed)
}

/*
Q4
You are tasked with creating a grade calculator program in GoLang.
The program should prompt the user to input their scores in three subjects:Math, Science, and English.
Based on these scores, calculate the average grade(assuming each subject is equally weighted)
and display the corresponding letter grade (A, B, C, D, or F)using control flow.
*/
func q4() {
	var (
		name    string
		math    float32
		science float32
		english float32
		average float32
		grade   string
	)

	fmt.Println("Enter the details of the student:")
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter score for Math: ")
	fmt.Scan(&math)
	fmt.Print("Enter score for Science: ")
	fmt.Scan(&science)
	fmt.Print("Enter score for English: ")
	fmt.Scan(&english)

	average = (math + science + english) / 3

	switch {
	case average >= 90:
		grade = "A"
	case average >= 80:
		grade = "B"
	case average >= 60:
		grade = "C"
	case average >= 50:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Grade scored by %s is %s\n", name, grade)
}

/*
Q5
You want to build a simple interest calculator in GoLang.
The program should ask the user to input multiple sets of data, each containing the principal amount,
the annual interest rate, and the number of years for which the interest is to be calculated.
Use arrays to store the input data for each set, calculate the simple interest for each set using the formula:
Simple Interest = (Principal Amount * Annual Interest Rate * Number of Years) / 100, and display the results.
*/
func q5() {

	const numSets = 3
	var principal [numSets]float64
	var interestRate [numSets]float64
	var years [numSets]float64
	var simpleInterest [numSets]float64

	for i := 0; i < numSets; i++ {
		fmt.Printf("Enter details for Set %d:\n", i+1)
		fmt.Print("Principal Amount: ")
		fmt.Scan(&principal[i])
		fmt.Print("Annual Interest Rate: ")
		fmt.Scan(&interestRate[i])
		fmt.Print("Number of Years: ")
		fmt.Scan(&years[i])

		simpleInterest[i] = (principal[i] * interestRate[i] * years[i]) / 100
	}

	fmt.Println("\nResults:")
	for i := 0; i < numSets; i++ {
		fmt.Printf("Set %d - Simple Interest: %.2f\n", i+1, simpleInterest[i])
	}
}

/*
Q6
Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants:/
*/
const maxStockQuantity = 200

type Product struct {
	Barcode  string
	Name     string
	Quantity int
}

func q6() {
	products := []Product{
		{Barcode: "123456", Name: "Apples", Quantity: 150},
		{Barcode: "789012", Name: "Bread", Quantity: 180},
		{Barcode: "345678", Name: "Milk", Quantity: 220},
	}
	for {
		fmt.Println("\nSupermarket Inventory Management")
		fmt.Println("1. Display Inventory")
		fmt.Println("2. Add Stock")
		fmt.Println("3. Sell Product")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayInventory(products)
		case 2:
			addStock(&products)
		case 3:
			sellProduct(&products)
		case 4:
			fmt.Println("Exiting. Thank you for using our system!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func displayInventory(products []Product) {
	fmt.Println("\nCurrent Inventory:")
	for _, p := range products {
		fmt.Printf("Barcode: %s | Product: %s | Quantity: %d\n", p.Barcode, p.Name, p.Quantity)
	}
}

func addStock(products *[]Product) {
	var barcode string
	var quantity int
	fmt.Print("Enter product barcode to add stock: ")
	fmt.Scan(&barcode)
	fmt.Print("Enter quantity to add: ")
	fmt.Scan(&quantity)

	for i, p := range *products {
		if p.Barcode == barcode {
			// Check if stock exceeds the maximum allowed
			if p.Quantity+quantity > maxStockQuantity {
				fmt.Printf("Cannot add more than %d units. Stock limit reached.\n", maxStockQuantity)
				return
			}
			(*products)[i].Quantity += quantity
			fmt.Printf("%d units added to %s.\n", quantity, p.Name)
			return
		}
	}
	fmt.Println("Product not found.")
}

func sellProduct(products *[]Product) {
	var barcode string
	var quantity int
	fmt.Print("Enter product barcode to sell: ")
	fmt.Scan(&barcode)
	fmt.Print("Enter quantity to sell: ")
	fmt.Scan(&quantity)

	for i, p := range *products {
		if p.Barcode == barcode {
			if p.Quantity >= quantity {
				(*products)[i].Quantity -= quantity
				fmt.Printf("%d units of %s sold.\n", quantity, p.Name)
			} else {
				fmt.Println("Insufficient stock.")
			}
			return
		}
	}
	fmt.Println("Product not found.")
}

/*
Q6.2
Calculate the factorial of a number.
*/
func factorial(number int64) int64 {
	if number == 1 {
		return 1
	}
	factorialOfNumber := number * factorial(number-1)
	return factorialOfNumber
}

func q6_2() {
	var number int64
	fmt.Println("Enter the number to find it's factorial")
	fmt.Scan(&number)
	fmt.Printf("The factorial of %d is %d.\n", number, factorial(number))
}
