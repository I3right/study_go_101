package main

import "fmt"

func reviewMovie(name string, rating float64) {
	fmt.Printf("I review %s and it's rating is %f \n", name, rating)
}

func findMovieName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: endgame"
	case "tt1825683":
		return "Black panther"
	}
	return "not found"
}

func byTicket(movie string) {
	fmt.Printf("I bought tickets to %s \n", movie)
}

func main() {
	movieName := findMovieName("tt4154796")
	fmt.Println(movieName)

	reviewMovie(movieName, 8.4)
	byTicket(movieName)
}
