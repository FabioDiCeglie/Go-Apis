package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "math/rand"
// "strconv"

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(){
}

func createMovie(){
}

func updateMovie(){
}

func deleteMovie(){
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ ID: "1", Isbn: "123456", Title: "First Movie", Director : &Director{FirstName: "fabio", LastName: "dc"} })
	movies = append(movies, Movie{ ID: "2", Isbn: "78910", Title: "Second Movie", Director : &Director{FirstName: "fabio", LastName: "dc"} })
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("8000", r))
}