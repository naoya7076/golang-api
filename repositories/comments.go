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

func SelectCommentList(db sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return []models.Comment{}, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		if err != nil {
			return []models.Comment{}, err
		} else {
			commentArray = append(commentArray, comment)
		}
	}
	return commentArray, nil
}
