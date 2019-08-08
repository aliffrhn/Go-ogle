package main

import "fmt"

func main(){
	
	var nilai = 9
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