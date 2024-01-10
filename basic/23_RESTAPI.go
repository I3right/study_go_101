/*
* REST API : return JSON
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	TiTle       string  `json:"title"`
	Year        string  `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []Movie

// - 											RESPONSE								REQUEST
func moviesHandler(w http.ResponseWriter, req *http.Request) {
	// receive request metohd
	method := req.Method

	if method == "POST" {
		//* read BODY
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "Error : %v", err)
			return
		}

		//	Read request body as String
		fmt.Println(string(body))

		//	Declare variable receive req body
		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			// handle errror when BAD REQUEST(format Body is not matching)
			w.WriteHeader(http.StatusBadRequest) //! 400
			fmt.Fprintf(w, "Error %s", err)
			return
		}

		movies = append(movies, t)

		fmt.Fprintf(w, "hello %s created movies", method)
		return
	}

	// fmt.Fprintf(w, "hello %v movies", movies)	//! <--- return like this is STRUC not JSON
	b, err := json.Marshal(movies) // convert STRUC -> JSON
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //! 500
		fmt.Fprintf(w, "Error %s", err)
		return
	}

	//* ----- RETURN VALUE

	// fmt.Fprintf(w, "%s", string(b))	//! %s is still return string that look like JSON
	// //*		^ Fprintf: write string in format printF and return value by that variable declare on function

	// w.Write(b) //! This Also return string not JSON

	//* SET Return as JSON
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.Write(b)
}

func main() {
	// Function in GO is first class value => able to pass function
	http.HandleFunc("/movies", moviesHandler)
	//?								PATH				FUNCTION

	//*		ListenAndServe: run on this url
	err := http.ListenAndServe("localhost:2565", nil)
	// if error log that error
	log.Fatal(err)
}
