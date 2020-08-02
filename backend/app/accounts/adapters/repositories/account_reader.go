package repositories

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/app/accounts/adapters/repositories/entities"
	"chat/app/dataproviders/sql"
	"chat/core"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := core.Injector.Provide(newAccountReader)
	if err != nil {
		log.Println("Error providing AccountReader instance:", err)
		panic(err)
	}
}

type accountReader struct {
	db *gorm.DB
}

func newAccountReader(conn sql.Connection) ports.AccountReader {
	db := conn.GetDatabase()
	return &accountReader{db}
}

func (u *accountReader) FindByID(id string) *accounts.Account {
	var account entities.Account
	u.db.Where("id = ?", id).First(&account)
	return account.ToDomain()
}

func (u *accountReader) FindByEmail(email string) *accounts.Account {
	var account entities.Account
	u.db.Where("email = ?", email).First(&account)
	return account.ToDomain()
}
