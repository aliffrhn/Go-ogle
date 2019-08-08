package main

import "fmt"

func main(){
	
	tampilkan_pesan1()
	fmt.Println(tampilkan_pesan2())
	// tampilkan_pesan3()
}

func tampilkan_pesan1(){
	fmt.Println("Pesan Diterima 1")
}

func tampilkan_pesan2()string{
	return "Pesan Diterima 2"
}