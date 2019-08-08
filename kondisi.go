package main

import "fmt"

func main(){
	
	// Define Nilai
	var nilai = 9
	
	// Condition
	if nilai == 10{
		fmt.Println("Lulus Dengan Baik")
	}else if nilai >= 7{
		fmt.Println("Lulus")
	}else if nilai == 6{
		fmt.Println("Belum Lulus")
	}else {
		fmt.Println("Belajar Lagi")
	}
}