package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
)

var (
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment1",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article",
		UserName:    "naoya7076",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
)

var (
	Article2 = Article{
		ID:        1,
		Title:     "first article",
		Contents:  "This is the test article",
		UserName:  "naoya7076",
		NiceNum:   1,
		CreatedAt: time.Now(),
	}
)
