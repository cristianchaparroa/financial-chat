package initializer

import (
	"chat/app/accounts/adapters/repositories/entities"
	"chat/app/dataproviders/sql"
	"chat/core"
)

func Migrate() {
	var conn sql.Connection
	invokeFunc := func(connection sql.Connection) {
		conn = connection
	}

	err := core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	db := conn.GetDatabase()

	db.AutoMigrate(entities.Account{})
}
