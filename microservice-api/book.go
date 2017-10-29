package main

import (
	"encoding/json"
	"net/http"
)

// book with name, author and isbn
type Book struct {
	Title string `json: "title"`
	Author string `json: "author"`
	ISBN string `json: "isbn"`
}

func (b Book) toJSON() []byte  {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func fromJSON(data []byte) Book  {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil{
		panic(err)
	}
	return book
}
var Books = []Book {
	Book{Title:"D", Author:"SDC", ISBN:"DFV"},
	Book{Title:"D", Author:"SDC", ISBN:"DFV"},
}
func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}