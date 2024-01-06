package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func main() {
	//? Struc -> JSON
	// c := Course{
	// 	Name:       "basic go",
	// 	Level:      "normal",
	// 	Views:      7239,
	// 	Instructor: "Anuchito",
	// 	FullPrice:  9999,
	// }

	// b, err := json.Marshal(c)
	// // marsgal -> sorting value and return byteSlice, err
	// //* result is JSON

	// fmt.Printf("Type : %T \n", b)
	// fmt.Printf("byte : %v \n", b)
	// fmt.Printf("string : %s \n", b)
	// fmt.Println("err", err)

	//? JSON -> Struc
	// data := []byte(`{
	// 	"name":	"basic go",
	// 	"level":	"normal",
	// 	"views":	7239,
	// 	"instructor":	"Anuchito",
	// 	"full_price":	9999
	// }`)

	// var c2 Course
	// err2 := json.Unmarshal(data, &c2)
	// fmt.Printf("%#v \n", c2)
	// fmt.Printf("%#v \n", c2.Instructor)
	// fmt.Println(err2)

	workshop()
}

type movie struct {
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	IsSuperHero bool     `json:"isSuperHero"`
}

func workshop() {
	// JSON -> struc -> movie
	body := `[
		{
			"imdbID":	"tt4154796",
			"title":	"Avengers: Endgame",
			"year":	2019,
			"rating": 8.4,
			"genres": ["Action","Drama"],
			"isSuperHero":	true
		},
		{
			"imdbID":	"tt4154756",
			"title":	"Avengers: Infinity war",
			"year":	2018,
			"rating": 9.2,
			"genres": ["Action","Sci-Fi"],
			"isSuperHero":	true
		}
	]`

	ms := []movie{}

	err := json.Unmarshal([]byte(body), &ms)
	if err != nil {
		fmt.Println("Error", err)
	}

	// fmt.Printf("%#v \n", ms)
	fmt.Println(ms)
}
