package entities

import (
	"chat/accounts"
	"chat/core/entities"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	*entities.Base
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// BeforeCreate will set the followings fields:
// - ID with uuid generated randomly.
// - Password: it will be save as a hash.
func (a *Account) BeforeCreate(scope *gorm.Scope) error {
	id := a.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}

	password := a.getPasswordHash()
	err = scope.SetColumn("Password", password)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) getPasswordHash() string {
	pass := []byte(a.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string
	return string(hash)
}

func NewFromDomain(u *accounts.Account) *Account {
	return &Account{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *Account) ToDomain() *accounts.Account {
	if u == nil {
		return nil
	}
	return &accounts.Account{
		ID:        u.ID,
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
