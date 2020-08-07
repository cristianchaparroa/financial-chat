package sql

import (
	"chat/core"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	err := core.Injector.Provide(NewMySQLConnection)
	core.CheckInjection(err, "MySQLConnection")
}

// getURL retrieves the URL to connection to SQL database.
func getURL(params ...string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", params[0], params[1], params[2], params[3], params[4])
}

// MySQLConnection represents a connection for MySQL database
type MySQLConnection struct {
	db *gorm.DB
}

// NewMySQLConnection retrieves a sql connection to MySQL server
func NewMySQLConnection() (Connection, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := getURL(dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	log.Println(url)

	db, err := gorm.Open("mysql", url)
	db.LogMode(true)

	if err != nil {
		return nil, err
	}

	return &MySQLConnection{db: db}, nil
}

// GetDatabase retrieve a MySQL connection
func (c *MySQLConnection) GetDatabase() *gorm.DB {
	return c.db
}
