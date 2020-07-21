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
	err := core.Injector.Provide(newUserWriter)
	if err != nil {
		log.Println("Error providing UserWriter instance:", err)
		panic(err)
	}
}

type userWriter struct {
	db *gorm.DB
}

func newUserWriter(conn sql.Connection) ports.UserWriter {
	db := conn.GetDatabase()
	return &userWriter{db}
}

func (u *userWriter) Create(user *chat.User) *chat.User {
	entity := entities.NewFromDomain(user)
	u.db.Model(entities.User{}).Save(&entity)
	return entity.ToDomain()
}
