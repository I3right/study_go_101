/*
*	DEFER	: Do later			[Stack: First in Last out]
! Before End function


? USE FOR ?
	- closing find: to make sure that find is going to close
*/

package main

import "fmt"

func main() {
	// No Defer
	// fmt.Println("World")
	// fmt.Println("Hello")

	// with Defer
	// defer fmt.Println("world")
	// fmt.Println("Hello")

	// defer fmt.Println("A")
	// defer fmt.Println("B")
	// defer fmt.Println("C")
	// defer fmt.Println("D")
	// fmt.Println("TEST")

	//! BEWARE SCOPE [Stack: First in Last out]
	for i := 0; i < 10; i++ {
		// defer fmt.Println("i", i)
		// defer fmt.Println("----------")
		// defer func() {
		// 	fmt.Println("i in F defer", i)
		// 	// ^ I going to be 10 all time cuz defer will active BF exit func so i will always be 10
		// 	// If want to be a number must be recieve from props
		// }()
		defer func(n int) {
			fmt.Println("n in F defer", n)
			// ^ I going to be 10 all time cuz defer will active BF exit func so i will always be 10
			// If want to be a number must be recieve from props
		}(i)
	}
	fmt.Println("DONE")
}
