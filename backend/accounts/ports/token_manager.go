package ports

import "chat/accounts"

type TokenManager interface {
	Generate(acc *accounts.Account) (string, error)
}
