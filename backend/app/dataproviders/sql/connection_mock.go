package sql

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

// MockConnection is an usefull connection for unit testing
type MockConnection struct {
	db *gorm.DB
}

// NewMockConnection setup a mock connection.
func NewMockConnection() Connection {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_mock")

	return &MockConnection{db: db}
}

// GetDatabase retrieves a mock connection.
func (c *MockConnection) GetDatabase() *gorm.DB {
	return c.db
}
