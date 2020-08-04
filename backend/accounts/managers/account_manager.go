package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/core"
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	err := core.Injector.Provide(newAccountManager)
	if err != nil {
		log.Println("Error providing AccountManager instance:", err)
		panic(err)
	}
}

type accountManager struct {
	writer ports.AccountWriter
	reader ports.AccountReader
}

func newAccountManager(w ports.AccountWriter, r ports.AccountReader) ports.AccountManager {
	return &accountManager{w, r}
}

func (m *accountManager) Create(u *accounts.Account) (*accounts.Account, error) {
	return m.writer.Create(u)
}

func (m *accountManager) GetByID(ID string) *accounts.Account {
	return m.reader.FindByID(ID)
}

func (m *accountManager) Login(acc *accounts.Account) (*accounts.Account, error) {
	u := m.reader.FindByEmail(acc.Email)

	if u == nil {
		return nil, errors.New(UserNotFound)
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(acc.Password))

	if err != nil {
		return nil, errors.New(UserNotAuthenticated)
	}
	return u, nil
}
