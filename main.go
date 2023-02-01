package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naoya7076/golang-api/handlers"
)

func main() {
	con := controllers.NewMyAppController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.CommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
