package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naoya7076/golang-api/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.CommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}
