package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/core"
)

func init() {
	err := core.Injector.Provide(newRegisterManager)
	core.CheckInjection(err, "RegisterManager")
}

type registerManager struct {
	manager ports.AccountManager
	token   ports.TokenManager
}

func newRegisterManager(manager ports.AccountManager, token ports.TokenManager) ports.RegisterManager {
	return &registerManager{manager: manager, token: token}
}

func (m *registerManager) Register(acc *accounts.Account) (*accounts.Account, string, error) {

	acc, err := m.manager.Create(acc)

	if err != nil {
		return nil, "", err
	}

	token, err := m.token.Generate(acc)
	if err != nil {
		return nil, "", err
	}

	return acc, token, nil
}
