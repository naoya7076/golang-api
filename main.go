package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/naoya7076/golang-api/handlers"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

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

	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment1",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article",
		UserName:    "naoya7076",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}
