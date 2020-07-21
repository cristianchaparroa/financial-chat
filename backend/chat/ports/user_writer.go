package ports

import "chat/chat"

// UserWriter is in charge to write all the data related to users.
type UserWriter interface {

	// Create write a new user in database
	Create(user *chat.User) *chat.User
}
