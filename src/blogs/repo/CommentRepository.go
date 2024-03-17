package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CommentRepository) GetById(id string) (model.Comment, error) {
	var blog model.Comment
	dbResult := repo.DatabaseConnection.Where("id = ?", id).First(&blog)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (repo *CommentRepository) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	dbResult := repo.DatabaseConnection.Find(&comments)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return comments, nil
}