package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Artcle... \n")
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func NiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
