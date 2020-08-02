package ports

import "chat/accounts"

// AccountWriter is in charge to write all the data related to accounts.
type AccountWriter interface {

	// Create write a new user in database
	Create(account *accounts.Account) (*accounts.Account, error)
}
