package repo

import (
	"blogs/model"
	"errors"

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

func (repo *CommentRepository) GetByBlogId(id string, page int, pageSize int) ([]model.Comment, int, error) {
	var comments []model.Comment
	var totalCount int64

	// First, get the total count of comments
	repo.DatabaseConnection.Model(&model.Comment{}).Where("blog_id = ?", id).Count(&totalCount)

	// Then, get the paginated list of comments
	result := repo.DatabaseConnection.Where("blog_id = ?", id).Find(&comments)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return comments, int(totalCount), nil
}

func (repo *CommentRepository) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	dbResult := repo.DatabaseConnection.Find(&comments)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return comments, nil
}

func (repo *CommentRepository) Create(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Create(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CommentRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", id).Delete(&model.Comment{})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("no comment found for deletion")
	}
	return nil
}

func (repo *CommentRepository) Update(comment *model.Comment) error {
	println(comment)
	dbResult := repo.DatabaseConnection.Model(&model.Comment{}).Where("id = ?", comment.ID).Updates(comment)
	if dbResult.Error != nil {
		println("error se pojavio u bazi")
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		println("Nista nije promenjeno")
		return errors.New("no key point found for update")
	}
	return nil
}
