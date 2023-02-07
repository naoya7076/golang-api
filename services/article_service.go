package services

import (
	"database/sql"
	"errors"

	"github.com/naoya7076/golang-api/apperrors"
	"github.com/naoya7076/golang-api/models"
	"github.com/naoya7076/golang-api/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article, err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList []models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList, err}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}

	if commentGetErr != nil {
		commentGetErr = apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, commentGetErr
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err := apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Article{}, err
	}
	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}
	return article, nil
}
