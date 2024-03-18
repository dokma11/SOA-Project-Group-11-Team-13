package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type VoteRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *VoteRepository) GetById(id string) (model.Vote, error) {
	var blog model.Vote
	dbResult := repo.DatabaseConnection.Where("id = ?", id).First(&blog)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (repo *VoteRepository) GetAll() ([]model.Vote, error) {
	var votes []model.Vote
	dbResult := repo.DatabaseConnection.Find(&votes)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return votes, nil
}