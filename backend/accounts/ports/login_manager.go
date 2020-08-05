package ports

import "chat/accounts"

type LoginManager interface {
	Login(acc *accounts.Account) (*accounts.Account, string, error)
}
