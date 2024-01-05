/*
Transfer Type value
* function type(value)
! if result after convert is over limit it turn into zero value
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	fmt.Printf("Type: %T, Val: %d \n", i, i)

	//* int -> float  => float64()
	var f float64 = float64(i)
	fmt.Printf("Type: %T, Val: %f \n", f, f)

	//* float -> unsign int  => float64()
	var u uint8 = uint8(f)
	fmt.Printf("Type: %T, Val: %d \n", u, u)

	//? STRING -> INT
	v := "72"
	//! in Go is not accept string => int
	// var i2 = int(v)
	// fmt.Printf("Type: %T, Val: %d \n", i2, i2)

	i2, err := strconv.Atoi(v)
	// Atoi => string to int and accept only string that int only
	// if sent int + string : error invalid syntax

	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("Type: %T, Val: %d \n", i2, i2)

	//? INT -> STRING
	i3 := 10
	s2 := strconv.Itoa(i3)
	fmt.Printf("Type: %T, Val: %s \n", s2, s2)

}
