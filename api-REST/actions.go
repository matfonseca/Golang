package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"IT", 2018, "RW"},
	Movie{"Joker", 2019, "AA"},
	Movie{"Invictus", 2005, "ZZ"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
func MovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	movieId := param["id"]
	fmt.Fprintf(w, "Selecciono la pelicula %s", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
