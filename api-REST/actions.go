package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"IT", 2018, "RW"},
		Movie{"Joker", 2019, "AA"},
		Movie{"Invictus", 2005, "ZZ"},
	}
	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	movieId := param["id"]
	fmt.Fprintf(w, "Selecciono la pelicula %s", movieId)
}
