package main

import "fmt"

func main(){

	// Define Array withour length
	var fruits = []string{"apple", "manggo", "melon", "pienaple"}
	
	// Check length fruits
	fmt.Println(len(fruits))

	// Call fruits
	fmt.Println(fruits)

	// Append data
	fruits = append(fruits, "Durian")

	// Check length fruits
	fmt.Println(len(fruits))

	// Call fruits1
	fmt.Println(fruits)
}