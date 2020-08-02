package ports

import (
	"chat/accounts"
)

// AccountManager to handle all business cases related to accounts.
type AccountManager interface {

	// Create save a new user in the system.
	Create(u *accounts.Account) (*accounts.Account, error)

	// GetByID retrieve a user given id.
	GetByID(ID string) *accounts.Account

	// It validates the credentials of users
	Login(u *accounts.Account) (*accounts.Account, error)
}
