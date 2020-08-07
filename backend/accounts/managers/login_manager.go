package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/core"
)

func init() {
	err := core.Injector.Provide(newLoginManager)
	core.CheckInjection(err, "LoginManager")
}

type loginManager struct {
	manager ports.AccountManager
	token   ports.TokenManager
}

func newLoginManager(manager ports.AccountManager, token ports.TokenManager) ports.LoginManager {
	return &loginManager{manager: manager, token: token}
}

func (m *loginManager) Login(acc *accounts.Account) (*accounts.Account, string, error) {

	acc, err := m.manager.Login(acc)

	if err != nil {
		return acc, "", err
	}

	t, err := m.token.Generate(acc)
	if err != nil {
		return acc, "", err
	}

	return acc, t, nil
}
