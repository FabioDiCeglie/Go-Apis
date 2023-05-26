package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FabioDiCeglie/Learning-Go/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting server at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
