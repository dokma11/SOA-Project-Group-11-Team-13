package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type VoteService struct {
	VoteRepository *repo.VoteRepository
}

func (service *VoteService) GetById(id string) (*model.Vote, error) {
	vote, err := service.VoteRepository.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("vote with id %s not found", id))
	}
	return &vote, nil
}

func (service *VoteService) GetAll() (*[]model.Vote, error) {
	votes, err := service.VoteRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no votes were found"))
	}
	return &votes, nil
}