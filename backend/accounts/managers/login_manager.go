package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/core"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := core.Injector.Provide(newLoginManager)
	if err != nil {
		log.Println("Error providing LoginManager instance:", err)
		panic(err)
	}
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
