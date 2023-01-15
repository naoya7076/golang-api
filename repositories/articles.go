package articles

import (
	"database/sql"
	"fmt"

	"github.com/naoya7076/golang-api/models"
)

const articleNumPerPage = 5

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ? 0, now()`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;`
	offset := (page - 1) * articleNumPerPage
	rows, err := db.Query(sqlStr, page, offset)
	if err != nil {
		fmt.Println(err)
		return []models.Article{}, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			return []models.Article{}, err
		} else {
			articleArray = append(articleArray, article)
		}
	}
	return articleArray, nil
}
