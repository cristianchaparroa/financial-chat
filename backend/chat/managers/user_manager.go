package managers

import (
	"chat/chat"
	"chat/chat/ports"
	"chat/core"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := core.Injector.Provide(newUserManager)
	if err != nil {
		log.Println("Error providing UserManager instance:", err)
		panic(err)
	}
}

type userManager struct {
	writer ports.UserWriter
	reader ports.UserReader
}

func newUserManager(w ports.UserWriter, r ports.UserReader) ports.UserManager {
	return &userManager{w, r}
}

func (m *userManager) Create(u *chat.User) *chat.User {
	return m.writer.Create(u)
}

func (m *userManager) GetByID(ID string) *chat.User {
	return m.reader.FinUserByID(ID)
}
