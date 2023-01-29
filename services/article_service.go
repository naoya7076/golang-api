package services

import (
	"database/sql"

	"github.com/naoya7076/golang-api/models"
	"github.com/naoya7076/golang-api/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	article, err := repositories.SelectArticleDetail(&sql.DB{}, articleID)
	if err != nil {
		return models.Article{}, err
	}
	commentList, err := repositories.SelectCommentList(sql.DB{}, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}
