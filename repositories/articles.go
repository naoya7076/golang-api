package articles

import (
	"database/sql"
	"fmt"

	"github.com/naoya7076/golang-api/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ? 0, now()`

	_, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return article, err
	}
	return article, nil
}
