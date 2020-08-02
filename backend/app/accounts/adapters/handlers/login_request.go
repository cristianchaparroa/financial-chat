package handlers

import (
	"chat/accounts"
	"chat/core/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password", validate:"required"`
}

func newLoginRequest(c *gin.Context) (*LoginRequest, error) {

	var req LoginRequest
	err := c.BindJSON(&req)

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *LoginRequest) IsValid() *handlers.Error {
	params := make([]handlers.InvalidParam, 0)

	validFields := r.HasValidFields()
	params = append(params, validFields...)

	if len(params) == 0 {
		return nil
	}

	error := handlers.NewErrorWithStatus(handlers.TitleInvalidRequestError, http.StatusBadRequest)
	return error
}

func (r *LoginRequest) HasValidFields() []handlers.InvalidParam {
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
		param := handlers.InvalidParam{
			Name:   handlers.ToSnakeCase(e.Field()),
			Reason: e.Translate(custom.Translator),
			Code:   handlers.RequiredFieldCode,
		}
		params = append(params, param)
	}
	return params
}

func (r *LoginRequest) GetUser() *accounts.Account {
	return &accounts.Account{
		Email:    r.Email,
		Password: r.Password,
	}
}
