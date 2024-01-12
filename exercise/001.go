package main

/*
Exercise 001:

!but are not a multiple of 5, between 2000 and 3200 (both included).
? The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Exercise 001")

	res := Ex001(2000, 3200)

	fmt.Println(res)

}

// Ex001 returns a slice of numbers
func Ex001(low, high int) string {
	var result []string

	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			result = append(result, strconv.Itoa(i))
		}
	}

	// fmt.Println("RS", strings.Join(result, ","))
	fmt.Println("RAW", result)
	return ""
}
