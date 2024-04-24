package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

type User struct {
	ID             int64    `json:"id"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Role           UserRole `json:"role"`
	ProfilePicture string   `json:"profilePicture"`
	IsActive       bool     `json:"isActive"`
}

func NewUser(id int64, username string, password string, role UserRole, profilePicture string, isActive bool) (*User, error) {
	user := &User{
		ID:             id,
		Username:       username,
		Password:       password,
		Role:           role,
		ProfilePicture: profilePicture,
		IsActive:       isActive,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Validate() error {
	if user.Username == "" {
		return errors.New("invalid username. Username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("invalid password. Password cannot be empty")
	}
	if user.Role < 0 || user.Role > 3 {
		return errors.New("invalid role")
	}
	if user.ProfilePicture == "" {
		return errors.New("invalid profile picture path. Profile picture path cannot be empty")
	}
	return nil
}

func (user *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(user)
}

func (user *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(user)
}

func (user *User) String() string {
	return fmt.Sprintf("User{ID: %d, Username: %s, Role: %d, IsActive: %t, ",
		user.ID, user.Username, user.Role, user.IsActive)
}
