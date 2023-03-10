package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naoya7076/golang-api/api/middlewares"
	"github.com/naoya7076/golang-api/controllers"
	"github.com/naoya7076/golang-api/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.NiceArticleHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.CommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)
	return r
}
