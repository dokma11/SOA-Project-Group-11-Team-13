package model

import "gorm.io/gorm"

type VoteType int
const (
	Downvote VoteType = iota
	Upvote
)

type Vote struct {
	gorm.Model
	ID int
	UserId int
	BlogId int
	Type VoteType
}

func (vote *Vote) BeforeCreate(scope *gorm.DB) error {
	if vote.ID == 0 {
		var maxID int
		if err := scope.Table("vote").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		vote.ID = maxID + 1
	}
	return nil
}