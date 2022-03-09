package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type anime struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allAnimes []anime

var animes = allAnimes{
	{
		ID:          "1",
		Title:       "Naruto",
		Description: "Um anime bom!",
	},
}

func indexHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(animes)
}

func createAnime(w http.ResponseWriter, r *http.Request) {

	var newAnime anime
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "The data of animes has been erros")
	}

	json.Unmarshal(reqBody, &newAnime)
	animes = append(animes, newAnime)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAnime)
}

func getAnime(w http.ResponseWriter, r *http.Request) {
	animeId := mux.Vars(r)["id"]

	for _, singleAnime := range animes {
		if singleAnime.ID == animeId {
			json.NewEncoder(w).Encode(singleAnime)
		}
	}
}

func updateAnime(w http.ResponseWriter, r *http.Request) {
	animeId := mux.Vars(r)["id"]

	var updatedAnime anime

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "This request has been errors in data")
	}

	json.Unmarshal(reqBody, &updatedAnime)

	for i, singleAnime := range animes {
		if singleAnime.ID == animeId {
			singleAnime.Title = updatedAnime.Title
			singleAnime.Description = updatedAnime.Description
			animes = append(animes[:i], singleAnime)
			json.NewEncoder(w).Encode(singleAnime)
		}
	}
}

func deleteAnime(w http.ResponseWriter, r *http.Request) {
	animeId := mux.Vars(r)["id"]

	for i, singleAnime := range animes {
		if singleAnime.ID == animeId {
			animes = append(animes, animes[i+1:]...)
			fmt.Fprint(w, "Anime:  %v has been deleted!", animeId)
		}
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHome).Methods("GET")
	router.HandleFunc("/anime", createAnime).Methods("POST")
	router.HandleFunc("/anime/{id}", getAnime).Methods("GET")
	router.HandleFunc("/anime/{id}", updateAnime).Methods("PATCH")
	router.HandleFunc("/anime/{id}", deleteAnime).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
