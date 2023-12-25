package main

import "fmt"

/*
	Zero Value (Initial Value)
	in Go when declare value but not assign it willl have a value

	String		->	""
	Int				->	0
	Float			->	0.00
	Boolean		->	false
	Rune			->	int		->	Nil (Null, No value)
*/

func main() {
	var title string
	var age int
	var price float64
	var check bool
	var r rune

	fmt.Printf("Title: %s \n", title)
	fmt.Printf("Age: %d \n", age)
	fmt.Printf("Price: %f \n", price)
	fmt.Printf("Check: %t \n", check)
	fmt.Printf("Rune: %c \n", r)
}
