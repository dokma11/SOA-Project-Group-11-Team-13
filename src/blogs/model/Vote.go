type VoteType int
const (
	Draft BlogStatus = iota
	Published
	Closed
	Active
	Famous
)

type Vote struct {
	gorm.Model
	ID int
	UserId int
	Type VoteType
}