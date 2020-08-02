package ports

import "chat/accounts"

// AccountReader is in charge to read all the data related to accounts.
type AccountReader interface {

	// FindByID retrieves an user that match with the given id.
	FindByID(id string) *accounts.Account

	FindByEmail(email string) *accounts.Account
}
