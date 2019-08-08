package main

import "fmt"

func main(){

	var harga_makanan = map[string]int{"ayam_goreng":20000, "sate_lalat":15000}
	
	// Show  
	fmt.Println(harga_makanan["ayam_goreng"])

	// Show Len
	fmt.Println(len(harga_makanan))
}