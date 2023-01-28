package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naoya7076/golang-api/models"
	"github.com/naoya7076/golang-api/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(*testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment",
	}
	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new comment id is expected %d but got %d", expectedCommentNum, newComment.CommentID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where comment_id = ? and article_id = ? and message = ?
		`
		testDB.Exec(sqlStr, comment.CommentID, comment.ArticleID, comment.Message)
	})
}
