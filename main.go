package main

import (
	"log"
	"net/http"

	"github.com/naoya7076/golang-api/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ListArticleHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.NiceArticleHandler)
	http.HandleFunc("/comment", handlers.CommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
