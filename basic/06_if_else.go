package main

import (
	"fmt"
	"math"
)

/*
	If else
	() is optional
	comparison 		== (double equal)
	And						&&
	Or						||
	More					>		>=
	Less					<		<=
*/

func main() {
	num := 34

	if num == 34 && (num > 30) {
		fmt.Println("Yes!! This num is Thirty Four")
		if num%2 == 0 {
			fmt.Println("Is EVEN")
		} else {
			fmt.Println("Is ODD")
		}
	}

	// 		SHORT IF
	lim := 225.0

	if v := math.Pow(10, 3); v < lim {
		fmt.Println("x power n is: ", v)
	} else {
		fmt.Printf("x power n is %g over limit %g. \n", v, lim)
	}

	// ! Waring v is the function variable scope -> v only knows in if

	workShop(8.4)
}

func workShop(num float64) {
	if num < 5 {
		fmt.Println("Disapointed 😞")
	} else if num >= 5 && num < 7 {
		fmt.Println("Normal 😐")
	} else if num >= 7 && num < 10 {
		fmt.Println("Good 🥰")
	} else {
		fmt.Println("TF 🤷🤷🤷🤷🤷")
	}
}
