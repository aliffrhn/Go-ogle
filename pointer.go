package main

import "fmt"

func main(){
	var number int =  10
	var address_number *int = &number

	fmt.Println("Number:",number)
	fmt.Println("Address Number:",address_number)
}