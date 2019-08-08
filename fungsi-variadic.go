package main

import "fmt"

func main(){

	var rata_rata = variadic(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Average:", rata_rata)
}

func variadic(numbers ...int)float64{

	var total int = 0

	for _, numbers := range numbers{
		total+= numbers
		fmt.Println(numbers)
	}

	var panjang int = len(numbers)

	return float64(total) / float64(panjang)

}