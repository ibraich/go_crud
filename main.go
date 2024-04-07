package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", ISBN: "111", Title: "John Wick", Director: &Director{Firstname: "Chad", Lastname: "Stahlski"}})
	movies = append(movies, Movie{ID: "2", ISBN: "112", Title: "Donnie Darko", Director: &Director{Firstname: "Richard", Lastname: "Kelly"}})
	movies = append(movies, Movie{ID: "3", ISBN: "113", Title: "Scott Pilgrim vs The World", Director: &Director{Firstname: "Edgar", Lastname: "Wright"}})
	r.HandleFunc("/movies", getMovies.Methods("GET"))
	r.HandleFunc("/movies{id}", getMovie.Methods("Get"))
	r.HandleFunc("/movies", createMovie.Methods("POST"))
	r.HandleFunc("/movies{id}", updateMovie.Methods("PUT"))
	r.HandleFunc("/movies{id}", deleteMovie.Methods("Delete"))
	fmt.Printf("Starting server at port:8000")
	log.Fatal(http.ListenAndServe(":8080", r))
}
