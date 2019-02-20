package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	ID string `json:"id"`
	Code string `json:"code"`
	Title string `json:"title"`
	Description string `json:"description"`
	Author *Author `json:"author"`

}

type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Init books
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through books and find with id

	for _, item := range books{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books   {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)

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

func deleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books   {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main(){
	r:= mux.NewRouter()

	//Data Example - @Todo implement Databases
	books = append(books,
		Book{ID:"1",Title:"Bumi",Author: &Author{FirstName:"Tere",LastName:"Liye"},Code:"123001"},
		Book{ID:"2",Title:"Bulan",Author: &Author{FirstName:"Tere",LastName:"Liye"},Code:"123002"},
		Book{ID:"3",Title:"Matahari",Author: &Author{FirstName:"Tere",LastName:"Liye"},Code:"123003"})
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))

}