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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "435656", Title: "erhab", Director: &Director{Firstname: "SEN", Lastname: "TENZ"}})
	movies = append(movies, Movie{ID: "2", Isbn: "35456", Title: "yalla bina", Director: &Director{Firstname: "AMINE", Lastname: "GOUERCH"}})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovies).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	router.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("STARTING THE SERVER AT PORT 8080 : \n ")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)

	for _, item := range movies {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}

}

func createMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-type", "application/json")
	//params
	param := mux.Vars(r)
	//loop over the movies , range
	//delete the movies with the i.d that you 've sent
	//add a enw movie - the movie that we send in the body of postman

	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}
