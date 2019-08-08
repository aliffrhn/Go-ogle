package main

import "fmt"

func main(){

	// Define Fruits With Arr
	var fruits = [4]string{"Apple", "Manggo", "Pienapple", "Cocoa"}
	
	// Get Length with len()
	fmt.Println("Length = ", len(fruits))

	// Get Fruits Before Update
	fmt.Println("Result Before Update = ", fruits)

	// Update
	fruits[1] = "Ayam"

	// After Update
	fmt.Println("Result After Update = ", fruits)
}