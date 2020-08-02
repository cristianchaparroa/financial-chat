package handlers

import "chat/accounts"

type RegisterUserResponse struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func newRegisterUserResponse(acc *accounts.Account, token string) *RegisterUserResponse {
	return &RegisterUserResponse{
		ID:        acc.ID,
		Token:     token,
		Email:     acc.Email,
		Name:      acc.Name,
		Password:  acc.Password,
		CreatedAt: acc.CreatedAt.Unix(),
		UpdatedAt: acc.UpdatedAt.Unix(),
	}
}
