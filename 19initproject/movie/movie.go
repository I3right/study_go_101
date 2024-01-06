package movie

import "fmt"

// ! lowercase => Private
// * Uppercase => Public

func init() {
	fmt.Println("init Movie")
}

func Review(name string, rating float64) {
	fmt.Printf("I review %s and it's rating is %f \n", name, rating)
}
