package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/books/getBookById", getBookById)
	http.ListenAndServe("", nil)
}

func getBookById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	writer.Header().Add("content-type", "application/json;charset=UTF-8")
	if id == "" {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusOK)
		book := Book{Name: "学习Go", Author: "诸葛亮", PubDate: time.Now()}
		json.NewEncoder(writer).Encode(book)
	}
}

type Book struct {
	Name    string
	Author  string
	PubDate time.Time
}
