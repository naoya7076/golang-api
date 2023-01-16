package repositories

import (
	"database/sql"
	"time"

	"github.com/naoya7076/golang-api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at)
		values (?, ?, now());
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message, newComment.CreatedAt = comment.ArticleID, comment.Message, time.Now()
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)
	return newComment, nil
}
