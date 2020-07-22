package chat

import "time"

// User contains the information related an user inside of room.
type User struct {
	ID        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
