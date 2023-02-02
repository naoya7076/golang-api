package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naoya7076/golang-api/controllers"
)

func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.NiceArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	return r
}
