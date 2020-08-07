package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInjection_WithPanic(t *testing.T) {
	serviceName := "BadService"
	err := errors.New("it's not possible to inject the service")
	assert.Panics(t, func() { CheckInjection(err, serviceName) }, "It expects a panic")
}

func TestCheckInjection_WithoutPanic(t *testing.T) {
	serviceName := "OKService"
	assert.NotPanics(t, func() { CheckInjection(nil, serviceName) }, "It doesn't expect a panic")
}
