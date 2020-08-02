package handlers

import (
	"chat/accounts/ports"
	"chat/core"
	"chat/core/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	err := core.Injector.Provide(newRegisterUserHandler)
	if err != nil {
		log.Println("Error providing RegisterUserHandler instance:", err)
		panic(err)
	}
}

type RegisterUserHandler struct {
	manager ports.RegisterManager
}

func newRegisterUserHandler(manager ports.RegisterManager) *RegisterUserHandler {
	return &RegisterUserHandler{manager: manager}
}

func (h *RegisterUserHandler) Create(c *gin.Context) {

	req, err := newLoginRequestFromContext(c)

	if err != nil {
		generateError(c, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()

	if error != nil {
		handlers.GenerateFullError(c, error)
		return
	}

	account := req.ToDomain()
	acc, token, err := h.manager.Register(account)

	if err != nil {
		generateError(c, http.StatusUnprocessableEntity, err)
		return
	}

	response := newRegisterUserResponse(acc, token)
	c.JSON(http.StatusCreated, response)
}
