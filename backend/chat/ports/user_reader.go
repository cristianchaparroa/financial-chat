package ports

import "chat/chat"

// UserReader is in charge to read all the data related to users.
type UserReader interface {

	// FinUserByID retrieves an user that match with the given id.
	FinUserByID(id string) *chat.User
}
