package repositories

import (
	"chat/app/chat/adapters/repositories/entities"
	"chat/app/dataproviders/sql"
	"chat/chat"
	"chat/chat/ports"
	"chat/core"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := core.Injector.Provide(newUserReader)
	if err != nil {
		log.Println("Error providing UserReader instance:", err)
		panic(err)
	}
}

type userReader struct {
	db *gorm.DB
}

func newUserReader(conn sql.Connection) ports.UserReader {
	db := conn.GetDatabase()
	return &userReader{db}
}

func (u *userReader) FinUserByID(id string) *chat.User {
	var user entities.User
	u.db.Where("id = ?", id).First(&user)
	return user.ToDomain()
}
