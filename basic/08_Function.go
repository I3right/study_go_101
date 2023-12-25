/*
	CONCEPT
	- First class function
	- Higher Order function
*/

package main

import (
	"fmt"
	"math"
)

/*
	- Go don't allow if not give agrument to parameter
	- Able to return multiple value but also variable to store that value
*/

func add(x float64, y float64) (float64, string) {
	fmt.Println("hello add pos ", x, y)
	return x + y, "test return"
}

// x and y is string
// if want to return multiple value have to tell type value that return
func swap(x, y string) (string, string) {
	return y, x
}

// naked return !not good for complex function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// ^ don't need := cuz declare what is return
	return x, y
}

// Function is firstclass function = function is value
// Able to store function in variable (Type = Function) OR
// that function declare as varaiable
// Signature of function = Type of function
// Benefit: able to change function of logic
var add2 = func(x, y float64) float64 {
	return x + y
}

func pitagorus(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func compute(fn func(float64, float64) float64) float64 {
	v := fn(99, 1)
	return v
}

/*
Higher Order function
function -> parameter & return is Function
*/
func increase() int {
	return 1
}
func current() int {
	return 2
}

// This is High order function
func adder() (func() int, func() int) {
	sum := 0
	// return increase, current
	return func() int {
			sum += 1
			return sum
		}, func() int {
			return sum
		}
}

func emote(rating float64) string {
	switch {
	case rating < 5:
		return "Disapointed 😞"
	case rating >= 5 && rating < 7:
		return "Normal 😐"
	case rating >= 7 && rating < 10:
		return "Good 🥰"
	default:
		return "The Fuck"
	}
}
func workShop() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}

func main() {
	// a, b := add(12, 24)
	// fmt.Println(a)
	// fmt.Println(b)

	// c, d := swap("hello", "world")
	// fmt.Println(c, d)

	// e := add2(42, 13)
	// fmt.Println(e)

	// f := compute(add2)
	// fmt.Println("function plus", f)

	// g := compute(pitagorus)
	// fmt.Println("function pitagorus", g)

	// inc, cur := adder()
	// inc()
	// fmt.Println("increase value", inc())
	// fmt.Println("current value", cur())

	workShop()
}
