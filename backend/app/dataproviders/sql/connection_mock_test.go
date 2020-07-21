package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockConnection(t *testing.T) {
	conn := NewMockConnection()
	db := conn.GetDatabase()

	isAlive := db.DB().Ping()
	assert.Nil(t, isAlive)
}
