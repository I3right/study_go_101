package main

import "fmt"

/*
* ---- POINTER
! Go Can't Manipulate memory on Low level
* Go pass value of price to change Price

? Go Always Pass by value
! Zero value of pointer => nil
*/

// Normal function that recive props as int not *int(pointer)
// price(variable) was change only in this function
func changePrice(price int) {
	price = price - 599
	fmt.Println("F.change value", price, &price)
}

// ! If want Go to change value must be change value at that pointer
// * Derefference
func changePricePointer(price *int) {
	*price = *price - 599
	fmt.Println("F.change pointer")
	fmt.Println("old pointer", price, "New pointer", &price)
}

// STRUCT
type course struct {
	name, instructor string
	price            int
}

// Normal function that recive							[VALUE]
func discount(c course) int {
	c.price = c.price - 599
	fmt.Println("Discount: ", c.price)
	return c.price
}

// Normal function that recive							[Pointer]
func discountPointer(c *course) int { // Function
	//(*c).price = (*c).price - 599 // * De-reference Go will simplifly by remove(*)
	c.price = c.price - 599
	fmt.Println("Discount: ", c.price)
	return c.price
}

func (c *course) discountPointer() int { // Method
	//(*c).price = (*c).price - 599 // * De-reference Go will simplifly by remove(*)
	c.price = c.price - 599
	fmt.Println("Discount: ", c.price)
	return c.price
}

func main() {
	var price int = 9999
	var addr *int = &price // Collect location of price on Memory
	// (Type *int) : Pointer of Integer

	// fmt.Println("1: ", price, addr)

	*addr = 9400 // Write at that memory Position

	// fmt.Println("2: ", price, addr)

	// v := *addr // Read at that memory
	// fmt.Println("v @pointer", v)

	// changePrice(price)			// * passsing value to this function
	// fmt.Println("after", price)

	// fmt.Println("before", price)
	// changePricePointer(addr) // * passing pointer to this function
	// fmt.Println("after", price)

	// ! -------------- Pointer with STRUCT
	// c := course{"Basic Go", "Anuchit", 9999}

	// * SENT VALUE
	// d := discount(c)
	// fmt.Println("Discount Price: ", d)
	// fmt.Println("price", c.price)

	// * SENT POINTER
	// e := discountPointer(&c)
	// fmt.Println("Discount Price: ", e)
	// fmt.Println("price", c.price)
	// if do this way Should declare c as Pointer not value
	// c := &course{"Basic Go", "Anuchit", 9999}
	// c := new(course)	//* This new function return Pointer value

	workShop()
}

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(r float64) {
	m.votes = append(m.votes, r)
}

func workShop() {
	av := &movie{"Avengers: Endgame", 2019, 8.4, []float64{7, 8, 9, 10, 6, 9, 9, 10, 8}, []string{"Action", "Drama"}, true}

	fmt.Println("before", av)
	av.addVote(8)
	fmt.Println("after", av)
}
