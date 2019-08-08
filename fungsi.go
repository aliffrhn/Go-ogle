package main

import "fmt"

func main(){
	
	tampilkan_pesan1()
	fmt.Println(tampilkan_pesan2())
	fmt.Println(tampilkan_pesan3(10,2))
	fmt.Println(tampilkan_pesan4(10,2))
}

func tampilkan_pesan1(){
	fmt.Println("Pesan Diterima 1")
}

func tampilkan_pesan2()string{
	return "Pesan Diterima 2"
}

func tampilkan_pesan3(x int, y int)int{
	var z = x * y
	return z
}

func tampilkan_pesan4(x int, y int)(int,int){
	var z = x * y
	var a = x / y
	return z,a
}