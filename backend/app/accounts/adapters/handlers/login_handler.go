package handlers

import (
	"chat/accounts/ports"
	"chat/core"
	"chat/core/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	err := core.Injector.Provide(newLoginHandler)
	core.CheckInjection(err, "LoginHandler")
}

type LoginHandler struct {
	manager ports.LoginManager
}

func newLoginHandler(manager ports.LoginManager) *LoginHandler {
	return &LoginHandler{manager: manager}
}

func (l *LoginHandler) Login(c *gin.Context) {

	req, err := newLoginRequest(c)
	if err != nil {
		generateError(c, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()
	if error != nil {
		handlers.GenerateFullError(c, error)
		return
	}

	acc := req.GetUser()
	acc, t, err := l.manager.Login(acc)

	if err != nil {
		generateError(c, http.StatusUnprocessableEntity, err)
		return
	}

	response := newRegisterUserResponse(acc, t)
	c.JSON(http.StatusOK, response)
}
