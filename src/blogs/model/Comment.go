package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int
	AuthorId  int
	BlogId    int
	CreatedAt time.Time
	UpdatedAt *time.Time //zvezda stoji jer ova vrednost moze biti null
	Text      string
}

func NewComment(authorID, blogID int, createdAt time.Time, text string) *Comment {
	return &Comment{
		AuthorId:  authorID,
		BlogId:    blogID,
		CreatedAt: createdAt,
		Text:      text,
	}
}

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	if comment.ID == 0 {
		var maxID int
		if err := scope.Table("comments").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		comment.ID = maxID + 1
	}
	return nil
}

func (comment *Comment) UpdateText(text string) {
	if text != "" {
		comment.Text = text
		now := time.Now()
		comment.UpdatedAt = &now
	}
}
