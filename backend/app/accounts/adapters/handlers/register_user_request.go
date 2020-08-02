package handlers

import (
	"chat/accounts"
	"chat/core/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// RegisterUserRequest contains the data to create a new user.
type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (r *RegisterUserRequest) IsValid() *handlers.Error {
	custom := handlers.NewValidator()
	err := custom.Validator.Struct(r)

	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)
	if len(errors) == 0 {
		return nil
	}

	params := make([]handlers.InvalidParam, 0)

	for _, e := range errors {
		param := handlers.InvalidParam{Name: handlers.ToSnakeCase(e.Field()), Reason: e.Translate(custom.Translator), Code: handlers.RequiredFieldCode}
		params = append(params, param)
	}
	error := handlers.NewErrorWithStatus(handlers.TitleInvalidRequestError, http.StatusBadRequest)
	error.InvalidParams = params
	return error
}

func (r *RegisterUserRequest) ToDomain() *accounts.Account {
	return &accounts.Account{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func newLoginRequestFromContext(c *gin.Context) (*RegisterUserRequest, error) {
	var request RegisterUserRequest

	err := c.BindJSON(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}
