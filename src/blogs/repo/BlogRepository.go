package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
	var blog model.Blog
	dbResult := repo.DatabaseConnection.Preload("Comments").Where("id = ?", id).First(&blog)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
	var blogs []model.Blog
	dbResult := repo.DatabaseConnection.Find(&blogs)
	if dbResult != nil {
		return nil, dbResult.Error
	}
	return blogs, nil
}