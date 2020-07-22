package ports

import (
	"chat/chat"
)

// UserManager to handle all business cases related to users.
type UserManager interface {

	// Create save a new user in the system.
	Create(u *chat.User) *chat.User

	// GetByID retrieve a user given id.
	GetByID(ID string) *chat.User
}
