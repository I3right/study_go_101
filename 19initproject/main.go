package main

import (
	"fmt"

	"github.com/Jirayu/cinema/movie"
	"github.com/Jirayu/cinema/ticket"
)

//* SEPERATE PACKAGE
//? Package movie
// func reviewMovie(name string, rating float64) {
// 	fmt.Printf("I review %s and it's rating is %f \n", name, rating)
// }

// func findMovieName(imdbID string) string {
// 	switch imdbID {
// 	case "tt4154796":
// 		return "Avengers: endgame"
// 	case "tt1825683":
// 		return "Black panther"
// 	}
// 	return "not found"
// }

//? Package Ticket
// func byTicket(movie string) {
// 	fmt.Printf("I bought tickets to %s \n", movie)
// }

//? main
// func main() {
// 	movieName := findMovieName("tt4154796")
// 	fmt.Println(movieName)

// 	reviewMovie(movieName, 8.4)
// 	byTicket(movieName)
// }

//* RESULT

func init() {
	fmt.Println("INIT FUNCTION will work when call package")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)

	// reviewMovie(movieName, 8.4)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
