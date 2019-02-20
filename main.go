package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func getBook(w http.ResponseWriter, r *http.Request){

}




func main(){
	r:= mux.NewRouter()

	r.Handle("/api/books",getBooks).Methods("GET")
	r.Handle("/api/books/{id}",getBooks).Methods("GET")
	r.Handle("/api/books",createBook).Methods("POST")
	r.Handle("/api/books/{id}",updateBook).Methods("PUT")
	r.Handle("/api/books/{id}",deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))

}