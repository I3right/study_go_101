package main

import "fmt"

/*
	* Array [length]TypeValue
	! Length is FIXED, value is smae type
	- declare init value var Arr [3]string = [3]string{"1","2","3"}
	- Similar to other languages
	- len(Array) : Function in go that return Length of array
	- on Declare length of array can use [...] = Length is depense on size of parameter of array
*/

func show(sk [3]string) {
	fmt.Printf("%#v\n", sk)
}

func workShop() {
	var geners = [3]string{"Action", "Adventure", "Fantasy"}

	fmt.Println("OUTPUT")
	fmt.Printf("genres[0]: %#v\n", geners[0])
	fmt.Printf("genres[1]: %#v\n", geners[1])
	fmt.Printf("genres[2]: %#v\n", geners[2])

	geners[1] = "Sci-Fi"
	fmt.Printf("genres[1]: %#v\n", geners[1])
}

func main() {
	// var skill string
	// fmt.Printf("%#v\n", skill)

	// // init value for array
	// var Arr [3]string = [3]string{"1", "2", "3"}
	// fmt.Printf("%#v\n", Arr)
	// fmt.Printf("%#v\n", len(Arr))

	// show(Arr)

	// small := [3]int{1, 2, 3}
	// fmt.Printf("%#v\n", small)

	workShop()

}
