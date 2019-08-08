package main

import "fmt"

func main(){

	var number int =  10
	var address_number *int = &number

	fmt.Println("Number:",number)
	fmt.Println("Address Number:",address_number)

	// 

	var tas string = "Hash"
	var address_tas *string = &tas
	
	fmt.Println("Tas:",tas)
	fmt.Println("Address Tas",address_tas)
}