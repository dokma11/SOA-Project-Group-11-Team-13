package model

import (
	"gorm.io/gorm"
)

type BlogStatus int

const (
	Draft BlogStatus = iota
	Published
	Closed
	Active
	Famous
)

type Blog struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Status      BlogStatus
	AuthorId    int
	// ClubId int - verovatno ne mora
	Comments []Comment `gorm:"foreignKey:BlogId"`
	Votes    []Vote    `gorm:"foreignKey:BlogId"`
	// VisibilityPolicy BlogVisibilityPolicy - mozda necemo morati komplikovati
}

func NewBlog(title, description string, status BlogStatus, authorId int) *Blog {
	return &Blog{
		Title:       title,
		Description: description,
		Status:      status,
		AuthorId:    authorId,
	}
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	if blog.ID == 0 {
		var maxID int
		if err := scope.Table("blogs").Select("COALESCE(MAX(id), 0)").Row().Scan(&maxID); err != nil {
			return err
		}
		blog.ID = maxID + 1
	}
	return nil
}

func (blog *Blog) UpvoteCount() int {
	count := 0
	for _, vote := range blog.Votes {
		if vote.Type == Upvote {
			count++
		}
	}
	return count
}

func (blog *Blog) DownvoteCount() int {
	count := 0
	for _, vote := range blog.Votes {
		if vote.Type == Downvote {
			count++
		}
	}
	return count
}

func (blog *Blog) VoteCount() int {
	count := 0
	for _, vote := range blog.Votes {
		if vote.Type == Upvote {
			count++
		} else {
			count--
		}
	}
	return count
}

func (blog *Blog) SetVote(userId int, voteType VoteType) {
	// Find existing vote
	var existingIndex = -1
	for i, v := range blog.Votes {
		if v.UserId == userId {
			existingIndex = i
			break
		}
	}

	// If existing vote found, remove it
	if existingIndex != -1 {
		blog.Votes = append(blog.Votes[:existingIndex], blog.Votes[existingIndex+1:]...)
	}

	// If existing vote found and type matches, return
	if existingIndex != -1 && blog.Votes[existingIndex].Type == voteType {
		blog.UpdateBlogStatus()
		return
	}

	// Add new vote
	blog.Votes = append(blog.Votes, NewVote(userId, blog.ID, voteType))
	blog.UpdateBlogStatus()
}

func (blog *Blog) UpdateBlogStatus() {
	voteCount := blog.VoteCount()
	commentCount := len(blog.Comments)

	switch {
	case voteCount < -2:
		blog.Status = Closed
	case voteCount >= 3 && commentCount >= 3:
		blog.Status = Famous
	case voteCount >= 2 && commentCount >= 2:
		blog.Status = Active
	default:
		blog.Status = Published
	}
}
