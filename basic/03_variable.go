package main

import "fmt"

// Variable
// infer type (don't need to declare tpye of variable)
// :=  works only function scope
var packageScope string = "outsside function"

func main() {
	// var msg string = "hello Gopher!!"
	// var msg2 bool = true
	// var msg3 float64 = 123.32

	// price, check := 23, false

	// fmt.Println("1", msg)
	// fmt.Println("2", msg2)
	// fmt.Println("3", msg3)
	// fmt.Println("4", packageScope)
	// fmt.Println("price", price)
	// fmt.Println("check", check)

	workShop()
}

func workShop() {
	// var title, years, rating, category, isSuperhero = "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	var title, category string = "Avengers: Endgame", "Sci-Fi"
	var years int = 2019
	var rating float32 = 8.4
	var isSuperhero bool = true

	fmt.Println("เรื่อง : ", title)
	fmt.Println("ปี : ", years)
	fmt.Println("เรตติ้ง : ", rating)
	fmt.Println("ประเภท : ", category)
	fmt.Println("ซุปเปอร์ฮีโร่ : ", isSuperhero)

	var RAWSTRING = `
	เรื่อง :  Avengers: Endgame
	ปี :  2019
	เรตติ้ง :  8.4
	ประเภท :  Sci-Fi
	ซุปเปอร์ฮีโร่ :  true
	`

	print(RAWSTRING)
}

// Scope group by { }
// static conly type language
// Go not accept declare variable in function and not use

// 		------- PRIMARY TYPE -------
// bool 			=> Boolean
// string 		=> String
// int 				=> Integer	(Go will select by it self)
// uint   		=> Unsign Ineger ( positive number )
// byte   		=> alias for uint8 , same as uint8
// rune				=>
// float32		=>
// complex64	=>
