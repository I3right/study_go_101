// 		---------		FORMAT TEXT		-----------

package main

import "fmt"

func main() {

	// var r rune = 'A'
	// // rune => alias type of int32. collect Text  as number

	// fmt.Println("r", r)
	// fmt.Printf("r: %c\n", r)

	// title, years, rating, category, isSuperhero := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true

	// fmt.Printf("tpye : %T -- เรื่อง : %s \n", title, title)
	// fmt.Printf("ปี : %d \n", years)
	// fmt.Printf("เรตติ้ง : %.2f \n", rating)
	// fmt.Printf("ประเภท : %#v \n", category)
	// fmt.Printf("ซุปเปอร์ฮีโร่ : %t \n", isSuperhero)

	/*
		Printf Format

		%c	->	print as Text
		%s	->	print as String
		%d	->	print as Decimal number
		%f	->	print as float number  (%.2f) float with 2 digit
		%t	->	print as Boolean
		%#v	->	print the value
		%T	->	print type of value

		also add \n to print as new line
		printf only print not go new lines by it self
	*/

	workShop()
}

func workShop() {
	title, years, rating, category, isSuperhero := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	icon := '★'

	fmt.Printf("เรื่อง : %s \n", title)
	fmt.Printf("ปี : %d \n", years)
	fmt.Printf("เรตติ้ง : %.2f \n", rating)
	fmt.Printf("ประเภท : %#v \n", category)
	fmt.Printf("ซุปเปอร์ฮีโร่ : %t \n", isSuperhero)
	fmt.Printf("เรื่องโปรด : %c \n", icon)
}
