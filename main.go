package main
import(
	"github.com/gorilla/mux"
	"fmt"
)
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
type Movie struct {
	Id       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func main{
	r:=mux.newRouter()
	r.HandleFunc("/movies"getMovies.Methods("GET"))
	r.HandleFunc("/movies{id}"getMovie.Methods("Get"))
	r.HandleFunc("/movies"createMovie.Methods("POST"))
	r.HandleFunc("/movies{id}"updateMovie.Methods("PUT"))
	r.HandleFunc("/movies{id}"deleteMovie.Methods("Delete"))
	

}
