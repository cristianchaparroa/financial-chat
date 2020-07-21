package repository

import (
	"chat/chat"
	"time"
)

type User struct {
	ID        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewFromDomain(u *chat.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToDomain() *chat.User {
	return &chat.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
