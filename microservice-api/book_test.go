package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T){
	book:= Book{Title:"a", Author:"b" , ISBN: "c"}
	json := book.toJSON()
	assert.Equal(t, `{"title": "a", "author": "b", "isbn": "c"}`,
		string(json), "book json marshelling is wrong")
}


func TestBookFromJSON(t *testing.T){
	json := []byte(`{"title": "a", "author": "b", "isbn": "c"}`)
	book := fromJSON(json)
	assert.Equal(t, Book{Title: "a", Author: "b", ISBN: "c"}, book, "book json unmarshelling is wrong")
}