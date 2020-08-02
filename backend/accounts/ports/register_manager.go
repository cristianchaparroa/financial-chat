package ports

import "chat/accounts"

type RegisterManager interface {
	Register(acc *accounts.Account) (*accounts.Account, string, error)
}
