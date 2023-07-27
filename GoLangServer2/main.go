package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

const port string = ":8000"

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request)    {}
func createMovie(w http.ResponseWriter, r *http.Request) {}
func updateMovie(w http.ResponseWriter, r *http.Request) {}
func deleteMovie(w http.ResponseWriter, r *http.Request) {}

func main() {
	movies = append(
		movies,
		Movie{
			ID:    "232",
			ISBN:  "234",
			Title: "Avengers",
			Director: &Director{
				FirstName: "Josh",
				LastName:  "Whedon",
			},
		},
		Movie{
			ID:    "123",
			ISBN:  "233",
			Title: "Oppenheimer",
			Director: &Director{
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
	)

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started at port %s \n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
