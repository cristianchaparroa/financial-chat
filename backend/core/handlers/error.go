package handlers

import (
	"encoding/json"
)

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
	Code   string `json:"code,omitempty"`
}

type Error struct {
	Status        int            `json:"status"`
	Title         string         `json:"title"`
	Type          string         `json:"type,omitempty"`
	InvalidParams []InvalidParam `json:"invalid_params,omitempty"`
	Code          string         `json:"code,omitempty"`
}

func NewErrorWithStatus(title string, status int) *Error {
	return &Error{Title: title, Status: status}
}

func NewErrorWithStatusAndCode(title string, code string, status int) *Error {
	return &Error{Title: title, Status: status, Code: code}
}

func (e *Error) String() string {
	if e == nil {
		return ""
	}
	b, _ := json.Marshal(e)
	return string(b)
}
