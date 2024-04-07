package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for ind, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:ind], movies[ind+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for ind, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:ind], movies[ind+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}

	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", ISBN: "111", Title: "John Wick", Director: &Director{Firstname: "Chad", Lastname: "Stahlski"}})
	movies = append(movies, Movie{ID: "2", ISBN: "112", Title: "Donnie Darko", Director: &Director{Firstname: "Richard", Lastname: "Kelly"}})
	movies = append(movies, Movie{ID: "3", ISBN: "113", Title: "Scott Pilgrim vs The World", Director: &Director{Firstname: "Edgar", Lastname: "Wright"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("Get")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("Delete")
	fmt.Printf("Starting server at port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
