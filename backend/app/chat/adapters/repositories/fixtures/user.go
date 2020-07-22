package fixtures

import (
	"chat/app/chat/adapters/repositories/entities"
	uuid "github.com/satori/go.uuid"
	"time"
)

func GetUser() *entities.User {
	now := time.Now()
	return &entities.User{
		ID:        uuid.NewV1().String(),
		Name:      "John Doe",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func GetUserEntity(id string) []map[string]interface{} {
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
