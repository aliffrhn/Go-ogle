package main

import "fmt"

func main(){

	// With For
	for i:=1;i<=10;i++{
		fmt.Println("Angka", i)
	}

	// Define Nilai
	var nilai = 1

	// With While
	for nilai<=10{
		fmt.Println("Nilai", nilai)
		nilai++
	}

	var value = 1
	// With Do While
	for{
		fmt.Println("Value", value)
		
		value++
		
		if value>10{
			break
		}
	}
}