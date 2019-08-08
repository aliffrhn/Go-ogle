package main

import "fmt"

func main(){
	
	var x1 pelajar
	x1.nama = "John"
	x1.kelas = 12
	x1.umur = 20

	fmt.Println("Nama:", x1.nama)
	fmt.Println("Kelas:", x1.kelas)
	fmt.Println("Umur:", x1.umur)
}

type pelajar struct{
	nama string
	kelas int
	umur int
}