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

// set the header content type to json and encode the movies into json
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get the ID in param and find the movie in slice and return it
func getMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type to json
	w.Header().Set("Content-Type", "applictaion/json")

	// get the url params from the request
	params := mux.Vars(r)

	// loop thorug the movies to find movie
	for _, movie := range movies {
		// found a movie and encode it into json and return to response
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// create a movie get data from the user and convert it into go struct
// and store into slice
func createMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type
	w.Header().Set("Content-Type", "applictaion/json")

	var movie Movie

	// read the data from the body and decode it into go struct
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// generate the random int for a unique id
	movie.ID = strconv.Itoa(rand.Intn(1000000))

	// append new movie to teh movie slice
	movies = append(movies, movie)

	// return the newly created movie
	json.NewEncoder(w).Encode(movie)
}

// update the movie from the id
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the json Content-Type
	w.Header().Set("Content-Type", "application/json")

	// get the params from the request
	params := mux.Vars(r)

	// loop through the movies
	for index, item := range movies {
		// find the movie
		if item.ID == params["id"] {
			// delete the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			// store the new movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// update the id
			movie.ID = params["id"]
			// append back to the movies slice
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// a movie func to delete a movie from the slice
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "applictaion/json")

	// get the params from the request
	params := mux.Vars(r)

	// loop thorug the movies
	for index, movie := range movies {
		// find the movie
		if movie.ID == params["id"] {
			// remove the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// return all the movies
	json.NewEncoder(w).Encode(movies)
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
