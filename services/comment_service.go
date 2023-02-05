package services

import (
	"github.com/naoya7076/golang-api/apperrors"
	"github.com/naoya7076/golang-api/models"
	"github.com/naoya7076/golang-api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}
