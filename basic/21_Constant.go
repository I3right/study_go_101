/*
* CONSTANT: constant variable
  - that can't reassign new value

* VALEDIC FUNCTIOn:
*/
package main

import "fmt"

const Pi = 3.14

type day int

func main() {
	const world = "test"
	fmt.Println("Hello", world)
	fmt.Println("Happy ", Pi, "Day")

	// This declare pattern also work for var => var(name = value)
	// const (
	// 	Sunday    = 1
	// 	Monday    = 2
	// 	Tuesday   = 3
	// 	Wednesday = 4
	// 	Thursday  = 5
	// 	Friday    = 6
	// 	Saturday  = 7
	// )//! ^ Declare like this is hard code

	//* iota : assign integer to variable start from 0
	const (
		_ = iota // dont need value 0
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday : ", Sunday)
	fmt.Println("Monday : ", Monday)
	fmt.Println("Tuesday : ", Tuesday)
	fmt.Println("Wednesday : ", Wednesday)
	fmt.Println("Thursday : ", Thursday)
	fmt.Println("Friday : ", Friday)
	fmt.Println("Saturday : ", Saturday)

	//? Valedic Function (...)
	//eg Println: valedictunction

	skill("js", "go", "python", "java")
}

func skill(xs ...string) {
	//*  ^ variable xs became slices to able recvie variable
	for _, s := range xs {
		fmt.Println("Skill: ", s)
	}
}
