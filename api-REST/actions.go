package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(status string) {
	this.Status = status
}

func (this *Message) setMessage(mes string) {
	this.Message = mes
}

func createMessage(status string, mes string) *Message {
	message := new(Message)
	message.setMessage(mes)
	message.setStatus(status)
	return message
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return session
}

func responseMovie(w http.ResponseWriter, status int, result Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func responseMovies(w http.ResponseWriter, status int, result []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}
func responseMessage(w http.ResponseWriter, status int, message Message) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

var colletion = getSession().DB("curso_go").C("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
func MovieList(w http.ResponseWriter, r *http.Request) {

	var result []Movie
	err := colletion.Find(nil).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	responseMovies(w, 200, result)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	movieId := param["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	result := Movie{}

	err := colletion.FindId(oid).One(&result)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	responseMovie(w, 200, result)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close() //cerrar la lectura del body

	err = colletion.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
	}
	responseMovie(w, 200, movie_data)
}
func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	movieId := param["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close() //cerrar la lectura del body

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}

	err = colletion.Update(document, change)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}

func MovieDelete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	movieId := param["id"]

	if !bson.IsObjectIdHex(movieId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieId)

	err := colletion.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	responseMessage(w, 200, *createMessage("success", "La pelicula con id: "+movieId+" ha sido borrada correctamente"))
}
