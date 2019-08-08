package main

import "fmt"

func main(){

	// Define Fruits With Arr
	var fruits = [4]string{"Apple", "Manggo", "Pienapple", "Cocoa"}
	
	// Get Length with len()
	fmt.Println("Length = ", len(fruits))

	// Get Fruits
	fmt.Println("Result Before Update = ", fruits)

	// Update
	fruits[1] = "Ayam"

	fmt.Println("Result After Update = ", fruits)
}