package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHander := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Artcle... \n")
	}

	listArticleHander := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	firstArticleHander := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article No.1\n")
	}

	niceArticleHander := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	commentHander := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHander)
	http.HandleFunc("/article/list", listArticleHander)
	http.HandleFunc("/article/1", firstArticleHander)
	http.HandleFunc("/article/nice", niceArticleHander)
	http.HandleFunc("/comment", commentHander)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
