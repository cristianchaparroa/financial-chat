package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	contentType            = "Content-Type"
	ProblemJSONContentType = "application/problem+json"
)

func GenerateError(c *gin.Context, e *Error) {
	log.WithFields(map[string]interface{}{"module": "errors"}).Error(e.String())

	if e.Status == http.StatusInternalServerError {
		AbortWithError(c, e.Status, e)
		return
	}

	c.Header(contentType, ProblemJSONContentType)
	c.JSON(e.Status, e)
}

func GenerateFullError(c *gin.Context, error *Error) {
	log.WithFields(map[string]interface{}{"module": "errors"}).Error(error.String())
	c.Header(contentType, ProblemJSONContentType)
	c.JSON(error.Status, error)
}

func AbortWithError(c *gin.Context, status int, e *Error) {
	c.AbortWithStatusJSON(status, e)
}

func NewDefaultServerError() *Error {
	return NewErrorWithStatusAndCode(
		InternalServerError,
		InternalServerErrorCode,
		http.StatusInternalServerError)
}
