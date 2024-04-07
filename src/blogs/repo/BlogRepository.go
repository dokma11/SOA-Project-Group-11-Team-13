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
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return blogs, nil
}

func (repo *BlogRepository) Save(blog *model.Blog) error {
	dbResult := repo.DatabaseConnection.Create(blog)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *BlogRepository) UpdateStatus(id string, status model.BlogStatus) (model.Blog, error) {
	var blog model.Blog
	dbResult := repo.DatabaseConnection.Where("id = ?", id).First(&blog)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	blog.Status = status
	updateDbResult := repo.DatabaseConnection.Model(&model.Blog{}).Where("id = ?", id).Updates(blog)
	if updateDbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}