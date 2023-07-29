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

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Success struct {
	Status  int   `json:"status"`
	Message Movie `json:"message"`
}
type SuccessMovies struct {
	Status  int     `json:"status"`
	Message []Movie `json:"message"`
}

const port string = ":8000"

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mssg := SuccessMovies{Status: 200, Message: movies}
	json.NewEncoder(w).Encode(mssg)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var err Error
	w.Header().Set("Content-Type", "application/json")
	for _, ele := range movies {
		if ele.ID == id {
			mssg := Success{Status: 200, Message: ele}
			json.NewEncoder(w).Encode(mssg)
			return
		}
	}
	err.Status = 404
	err.Message = "Movie not found"
	json.NewEncoder(w).Encode(err)

}
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
