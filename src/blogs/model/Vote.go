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