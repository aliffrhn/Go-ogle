package main

import "fmt"

func main(){
	// Positif
	var number_positif8 uint8 =  255 //0-255
	var number_positif16 uint16 = 65534 //0-65535
	var number_positif32 uint32 = 4000000 //0-4294967295
	var number_positif64 uint64 = 400000000000 //0-

	fmt.Println(number_positif8)
	fmt.Println(number_positif16)
	fmt.Println(number_positif32)
	fmt.Println(number_positif64)

	// Negatif
	var number_negatif int8 = -99 //-128 - 127
	fmt.Println(number_negatif)

	// Desimal
	var number_desimal = 2.55
	fmt.Println(number_desimal)

	// Boolean
	var isExist bool = false
	fmt.Println(isExist)
}