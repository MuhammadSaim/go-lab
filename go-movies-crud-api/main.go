package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// important struct to store the movies

// a struct to store the movie data
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// a struct to store the Director's data
type Director struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

// init create slice to store movies insted of in DB
var movies []Movie

// TODO: return all the movies
func getMovies(w http.ResponseWriter, r *http.Request) {
}

// TODO: return the specific movie with an ID
func getMovie(w http.ResponseWriter, r *http.Request) {
}

// TODO: create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
}

// TODO: udpate a movie with specific ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
}

// TODO: delete a movie with specific ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
}

// a main entry point of an applictaion
func main() {
	// init mux router
	r := mux.NewRouter()

	// appending some movies data to check api is working
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438345",
		Title: "Movie One",
		Director: &Director{
			Firstname: "Jhon",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "676878",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "James",
			Lastname:  "Cameron",
		},
	})

	// init the routes of an applictaion
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// starting the server
	fmt.Printf("Starting the server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
