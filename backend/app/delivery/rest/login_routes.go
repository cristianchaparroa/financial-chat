package rest

import (
	_ "chat/accounts/managers"
	_ "chat/app/accounts/adapters/repositories"
	_ "chat/app/dataproviders/sql"
)

func setupLoginRoutes(s *server) {
	login, err := loadLoginHandler()
	checkError(err)

	register, err := loadRegisterHandler()
	checkError(err)

	s.router.GET("/login", login.Login)
	s.router.POST("/register", register.Create)
}
