package sql

import "github.com/jinzhu/gorm"

// Connection allows to retrieve the gorm database connection.
type Connection interface {
	GetDatabase() *gorm.DB
}
