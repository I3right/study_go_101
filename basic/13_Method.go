/*
* ----- METHOD

 */

package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

// * this is normal function
func discount(c course) int {
	p := c.price - 599
	fmt.Println("Left: ", p)
	return p
}

// * method: is a function but is useonly of struc
// can change from function to method of the struc
// by puting RECIVER before method name
func (c course) discount(d int) int {
	p := c.price - d
	fmt.Println("Left: ", p)
	return p
}
func (c course) info() {
	fmt.Println("Name", c.name)
	fmt.Println("Instructor", c.instructor)
	fmt.Println("Price", c.price)
}

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func (m movie) info() {
	fmt.Printf("%s: (%d) - %.2f \n", m.title, m.year, m.rating)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t %s \n", g)
	}
}

func workShop() {
	ae := movie{
		title:       "Avengers Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	ae.info()
}

func main() {
	// c := course{"basic Go", "Anchit", 9999}
	// fmt.Printf("%#v \n", c)

	// d := discount(c) // normall of using function
	// fmt.Println("discounted price by function", d)

	// e := c.discount(899) // normall of using function
	// fmt.Println("Discount price by method: ", e)

	// c.info()

	workShop()
}
