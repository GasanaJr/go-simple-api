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

type Item struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Items)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Items {
		if item.ID == params["id"] {
			deletedItem := Items[index]
			Items = append(Items[:index], Items[index+1:]...)
			json.NewEncoder(w).Encode(deletedItem)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	Items = append(Items, Item{
		ID:    strconv.Itoa(rand.Int()),
		Isbn:  "1234567890123",
		Title: "Item 1",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	Items = append(Items, Item{
		ID:    strconv.Itoa(rand.Int()),
		Isbn:  "1234567890124",
		Title: "Item 2",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	r.HandleFunc("/Items", getItems).Methods("GET")
	// r.HandleFunc("/Items/{id}", getItem).Methods("GET")
	// r.HandleFunc("/Items", createItem).Methods("POST")
	// r.HandleFunc("/Items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/Items/{id}", deleteItem).Methods("DELETE")

	fmt.Printf("Server listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
