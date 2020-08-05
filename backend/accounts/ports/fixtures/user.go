package fixtures

import (
	"chat/accounts"
	uuid "github.com/satori/go.uuid"
	"time"
)

func GetAccount() *accounts.Account {
	now := time.Now()
	return &accounts.Account{
		ID:        uuid.NewV1().String(),
		Name:      "John Doe",
		Email:     "johndoe@gmail.com",
		Password:  "12345",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func GetAccountEntity(id string) []map[string]interface{} {
	now := time.Now()
	return []map[string]interface{}{
		{
			"id":         id,
			"name":       "John Doe",
			"created_at": now.String(),
			"updated_at": now.String(),
		},
	}
}
