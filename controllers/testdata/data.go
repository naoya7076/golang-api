package testdata

import "github.com/naoya7076/golang-api/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "naoya7076",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	models.Article{
		ID:          2,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "naoya7076",
		NiceNum:     2,
		CommentList: commentTestData,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
