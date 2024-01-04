/*
* ---- Interface
	* Acceptance of data that able to use in function

	- empty interface => don't declare any type to this variable
	* Go 1.8 (+ Any)

*/

package main

import "fmt"

func show(val int) {
	fmt.Println(val)
}

func show2(val any) {
	i, ok := val.(int)
	// ! Declare [I] as new variable to collect v and check is INT
	// * Declare [ok] as new variable is bool ant true if val is INT

	// use ok variable as condition to do correct logic

	if ok {
		i = i + 2
		fmt.Println("i", i)
	} else {
		fmt.Println("Your current value is ", i)
		fmt.Println("Please enter new value as Integer!")
	}

}

// Also can handle type of [Val] by using switch case of val
// Recive type of val and switch case int .... case string ...
func show3(val any) {
	switch t := val.(type) {
	case int:
		i := t + 2
		fmt.Println("Type Int: ", i)
	case string:
		s := t + " add word"
		fmt.Println("Type String: ", s)
	default:
		fmt.Printf("Didn't handle %T type \n", t)
	}
}

// Naming Interface: Go recommend that ended name with "er"
type promotion interface {
	discount() int // Method discount that not accept parameter and return int
}

type infoer interface {
	info()
}

type course struct{}

func (c course) discount() int {
	return 123
}

func (c course) info() {
}

func sale(val promotion) {
	// fmt.Println("Sale : ", val)
	fmt.Printf("SALE: %#v \n", val)

	// cuz interface promotion require method dicount
	// * Able to use method discount on function sale
	val.discount()

	// ! but not able to using other method
	// val.info()
}

func newShow(val infoer) {
	val.info()
}

// * Embeding interface => Combine interface
type presenter interface {
	promotion
	infoer
}

func summary(val presenter) {
	val.discount()
	val.info()
}

func main() {
	// old
	// var v interface{}
	// 1.8
	// var v any
	// v = 36
	// fmt.Printf("%T %#v \n", v, v)

	// v = "hello"
	// fmt.Printf("%T %#v \n", v, v)

	// var v2 any
	// v2 = 36
	// show(v2) 	=> 	This is not working because we declare v2 as ANY
	// ! we must Tell Type to function show - TYPE ASSERTION
	// show(v2.(int))

	// * if don't want to type assertion. had to modified func show from int to any
	// show2(v2)

	// ? What if we sent as String not Integer
	// v3 := 234.43
	// show2(v3) // is ERROR due to v3 is String not Int
	// show3(v3)

	// v4 := course{}
	// sale(v4)

	// newShow(v4)
	// summary(v4)

	workShop()
}

type voter interface {
	// addVote(float64)
	addVote(rating float64)
}

func vote(v voter, rating float64) {
	v.addVote(rating)
}

type movie struct {
	title       string
	year        int
	rating      float32
	vote        []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) {
	m.vote = append(m.vote, rating)
}

func workShop() {
	av := &movie{"Avengers: Endgame", 2019, 8.4, []float64{7, 8, 9, 10, 6, 9, 9, 10, 8}, []string{"Action", "Drama"}, true}

	vote(av, 8)
	fmt.Println("Votes: ", av)
}
