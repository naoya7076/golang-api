package repositories_test

import (
	"testing"

	"github.com/naoya7076/golang-api/models"
	"github.com/naoya7076/golang-api/repositories"
	"github.com/naoya7076/golang-api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "test",
		UserName: "naoya7076",
	}
	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d", expectedArticleNum, newArticle.ID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	if err = repositories.UpdateNiceNum(testDB, articleID); err != nil {
		t.Fatal(err)
	}
	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}
	t.Cleanup(func() {
		const sqlStr = `
			update articles set nice = nice - 1 where article_id = ?
		`
		testDB.Exec(sqlStr, articleID)
	})
}
