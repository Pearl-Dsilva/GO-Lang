package main

/*
Domain: General Store Management
Scenerio:
Lab Exercise:
1. Implement the concept of Variables,values and type.
2. Implement the concept of control flow.
3. Implement the concept of Array and Slice.
4. Implement the concept of Map and Structs.
5. Implement the concept of functions and error handling
6. Implement the concept of interface
*/
import (
	"fmt"
)

type Store struct {
	storeName string
	rating    int
}

var size int = 10

func main() {
	var choice int
	fmt.Println("Enter your choice")
	fmt.Println("1. Add toys to the inventory")
	fmt.Println("2. Get toys toys of certain cost range from inventory")
	fmt.Println("3. Add toys to the inventory")
	fmt.Scan(choice)

}
