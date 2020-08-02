package accounts

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Account contains the information related an user inside of room.
type Account struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// It saves the password using a hash function.
func (u *Account) SetHashPassword(password string) *Account {
	pass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	u.Password = string(hash)
	return u
}
