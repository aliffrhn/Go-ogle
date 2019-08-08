package main

import "fmt"

func main(){
	var number int =  10
	var address_number *int = &number

	fmt.Println(number)
	fmt.Println(address_number)
}