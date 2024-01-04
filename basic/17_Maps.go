/*
*	----	MAPS    = is >   SET
* DATA STRUCTURE

! Key cant Repeat

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// // slice
	// // var ss = []string{}

	// // maps		[key]	value = initail value
	// var m map[string]int = map[string]int{}
	// // fmt.Println("map init", m)

	// //* ADD VALUE
	// m["Answer"] = 42
	// m["Bright"] = 27
	// // fmt.Println("map Add", m)
	// // fmt.Printf("Pf Real value: \t %#v \n", m)

	// //? read value inside map
	// v := m["Answer"]
	// fmt.Println("value", v)

	// // no value to read
	// v2 := m["test"]
	// fmt.Println("value", v2)

	// //! DELETE value
	// delete(m, "Answer")
	// fmt.Printf("After Delete: \t %#v \n", m)

	// // No n return zero value of "Answer" => as int => return 0
	// // n := m["Answer"]
	// // fmt.Println("value n after delete", n)

	// // ? what if Answer have value but same as zero value ?
	// m["Answer"] = 0
	// n, isValue := m["Answer"]
	// // * ^ Go have 2 var that return bool that have value?
	// if isValue {
	// 	fmt.Println("Yes, have value")
	// 	fmt.Println("Value is: ", n)
	// } else {
	// 	fmt.Println("No, Dont have")
	// }

	// fmt.Println("-----------------------------------")

	// //! ---- ZERO Value ------
	// var newM map[string]int     // zero value is nill
	// newM = make(map[string]int) // zero value is emptay map map[]

	// if newM == nil {
	// 	fmt.Println("new Map is nill")
	// }

	// fmt.Printf("Zero value: %#v \n", newM)

	workShop()
}

func WordCount(words string) map[string]int {
	arrStr := strings.Fields(words)
	var wordMaps = make(map[string]int)

	for _, key := range arrStr {
		// mapValue, isValue := wordMaps[key]

		// if !isValue {
		// 	wordMaps[key] = 1
		// } else {
		// 	wordMaps[key] = mapValue + 1
		// }

		wordMaps[key] = wordMaps[key] + 1

	}

	return wordMaps
}

func workShop() {
	s := "If it looks like a duck swims like a duck and quacks like a duck then it probalbly is a duck"
	w := WordCount(s)
	fmt.Printf("%#v \n", w)
}
