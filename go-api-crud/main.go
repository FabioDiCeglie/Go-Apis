package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "encoding/json"
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

var movie []Movie

func main() {
	r := mux.NewRouter()

	// movies = append([])
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("8000", r))
}