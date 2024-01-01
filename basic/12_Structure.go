package main

import "fmt"

/*
	* Struc is a set of data
	! If dont add value the zero value is going to be as the primary type of stuc
*/

// This is STRUC
type course struct {
	name  string
	study string
	price float64
}

type movie struct {
	title       string
	years       float64
	rating      float64
	genres      []string
	isSuperHero bool
}

func workShop() {
	avengers := movie{"EndGame", 2019, 8.4, []string{"Action", "Drama"}, true}
	InfinityWars := movie{
		title:       "Infinity wars",
		years:       2018,
		rating:      8.4,
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}

	var mvs []movie
	mvs = append(mvs, avengers, InfinityWars)

	fmt.Println(mvs)

}

func main() {
	// name := "basic Go"
	// study := "Bright"
	// price := 55
	// fmt.Println("name: ", name, " study by : ", study, " on package price: ", price)

	// Declaration Stuc
	// * 1. Full
	// ! Don't need to declare all of prop in struc
	// c := course{
	// 	name:  "New journey",
	// 	study: "New lang",
	// 	price: 0,
	// }

	// c.study = "WTH"

	// fmt.Println("name", c.name)
	// fmt.Println("study on", c.study)
	// fmt.Println("price", c.price)

	// *. Minimal
	// c2 := course{
	// 	"Wth i'm doing", "idk", 5555,
	// }

	// fmt.Println("name", c2.name)
	// fmt.Println("study on", c2.study)
	// fmt.Println("price", c2.price)

	workShop()
}
