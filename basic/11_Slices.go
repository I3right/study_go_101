package main

/*
	* Slices Similar to ARRAY
		- Able to use metohd like Array

	! But length is Dynamic that we don't declare length we can add or remove value to slice
		* Append Return new slice | It DONT Mutate OLD VALUE

	? It can slices itself with [Start : End]
		!!! Zero value is nil (Null)

	Create slice with init value =>  method make

*/

import (
	"fmt"
)

func workShop() {
	// Make a slice Votes that combine XS and YS then slice the value from index 5 -> 8
	XS := []float64{4, 5, 7, 8, 3, 8, 0}
	YS := []float64{7, 2, 10, 9, 7}

	var vote []float64
	vote = append(XS, YS...) // => { XS, ... YS }

	vote = vote[5:9]
	fmt.Printf("Value %#v", vote)
}

func main() {

	// * Array
	// Array := [3]string{"test1", "What", "damn"}
	// fmt.Printf("%T => %#v \n", Array, Array)

	// * Slice
	// Slice := []string{"test1", "What", "damn"}
	// fmt.Printf("%T => %#v \n", Slice, Slice)

	// * Add value to slice
	// ! It DONT Mutate OLD VALUE
	// Slice := []string{"test1", "What", "damn"}
	// fmt.Printf("length: %d --- value: %#v \n", len(Slice), Slice)

	// Slice = append(Slice, "Add value") // can add multiple value by  , value , value , ...
	// fmt.Printf("length: %d --- value: %#v \n", len(Slice), Slice)

	// * slice method
	// ! It DONT Mutate OLD VALUE
	// sliceValue := Slice[0:2]
	// fmt.Println(sliceValue)
	// If want to slice that return to the End can use 2 ways
	// - declae length of slice and use on [start : length Vartiable]
	// - or shorthand Stat point or End Point  but dont tell the end value
	// 		[0:] 	to End point
	// 		[:l]	from start point
	// 		[:]		From start point to End point
	// l := len(sliceValue)
	// sliceEndFull := Slice[0:l]
	// fmt.Println("Slice to end Full", sliceEndFull)

	// sliceEndShort := Slice[1:]
	// fmt.Println("Slice to end Short", sliceEndShort)

	// * Make method
	// ArrM := make([]int, 3)
	// fmt.Printf("%#v", ArrM)

	workShop()
}
