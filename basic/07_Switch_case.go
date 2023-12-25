package main

import (
	"fmt"
)

/*
	Switch Case
	() is optional
	break is not necessary
	fallthrough fall to lower case
*/

func main() {
	today := "Saturday"

	switch today {
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
		fallthrough
	case "Monday", "Tuesday":
		fmt.Println("Today is Weekday T^T")
	default:
		fmt.Printf("Today is %v not a Weekend T^T \n", today)
	}

	// 		SHORT swirch
	switch day := "Saturday"; day {
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
		fallthrough
	case "Monday", "Tuesday":
		fmt.Println("Today is Weekday T^T")
	default:
		fmt.Printf("Today is %v not a Weekend T^T \n", today)
	}

	// ! Waring day is the function variable scope -> day only knows in switch case

	// Switch No expression
	num := 1
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0:
		fmt.Println("Positive number")
	default:
		fmt.Println("The Fuck")
	}

	workShop(9.4)
}

func workShop(num float64) {
	switch {
	case num < 5:
		fmt.Println("Disapointed 😞")
	case num >= 5 && num < 7:
		fmt.Println("Normal 😐")
	case num >= 7 && num < 10:
		fmt.Println("Good 🥰")
	default:
		fmt.Println("The Fuck")
	}
}
