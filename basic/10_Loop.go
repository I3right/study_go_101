/*
	* In Go. There is only loop is FOR.
	For loop that can do as For loop, Do...While, While Loop
	- Forever Loop : not declare condition -> use for handle at Concurrency

*/

package main

import "fmt"

func workShop() {
	var Arr = [3]string{"Action", "Adventure", "Fantasy"}

	fmt.Printf("Before for loop: %#v\n", Arr)
	for i := 0; i < len(Arr); i++ {
		Arr[i] = "Movie: " + Arr[i]
	}
	fmt.Printf("After for loop: %#v\n", Arr)

	for i := 0; i < len(Arr); i++ {
		// fmt.Println(Arr[i])
		fmt.Printf("genres[%d]: %s \n", i, Arr[i])
	}
}

func main() {
	// ! Normal For Loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// ! While Loop
	// sum := 1
	// for sum < 1000 { // if condition true: DO the Code
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// ! Forever
	// for {
	// 	// logic
	// }

	// var Arr = [3]string{"1", "2", "3"}
	// for i := 0; i < len(Arr); i++ {
	// 	fmt.Println("val", Arr[i])
	// }

	// ! Range
	// for _, val := range Arr {
	// 	// fmt.Println("i", i, "val", val)
	// 	fmt.Println("val", val)
	// }

	workShop()
}
