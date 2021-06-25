package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct{
	ID string `json:"id"`
	Ispn string `json:"ispn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

type Author struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Init books variable as a slice Book struct
var books []Book 

// Get all books
func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a book
func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books to find ID
	for _, v := range books {
		if v.ID == params["id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a book
func createBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book 
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update a book
func updateBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, v := range books {
		if v.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book 
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func deleteBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, v := range books {
		if v.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main(){
	// Init the router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{
		ID: "1",
		Ispn: "2333",
		Title: "Book A",
		Author: &Author{
			FirstName: "Rahul",
			LastName: "Nair",
		},
	})
	books = append(books, Book{
		ID: "2",
		Ispn: "2543",
		Title: "Book B",
		Author: &Author{
			FirstName: "Rohan",
			LastName: "Nair",
		},
	})

	// Router handler
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books",createBooks).Methods("POST")
	router.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))

}
