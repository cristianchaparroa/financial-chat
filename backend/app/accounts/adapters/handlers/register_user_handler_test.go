package handlers

import (
	"bytes"
	"chat/accounts/managers"
	accountFixtures "chat/accounts/ports/fixtures"
	"chat/accounts/ports/mocks"
	"chat/app/accounts/adapters/handlers/fixtures"
	testingHelper "chat/core/handlers/testing"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type registerUserHandlerSuite struct {
	suite.Suite
	registerManager *mocks.RegisterManager
	handler         *RegisterUserHandler
}

func TestRegisterUserHandlerSuitInit(t *testing.T) {
	suite.Run(t, new(registerUserHandlerSuite))
}

func (s *registerUserHandlerSuite) SetupSuite() {
	s.registerManager = &mocks.RegisterManager{}
	s.handler = newRegisterUserHandler(s.registerManager)
}

func (s *registerUserHandlerSuite) TestCreate_InvalidRequest() {
	data := fixtures.GetInvalidRegisterRequest()
	request := testingHelper.GetRequestBody(data)
	recorder := s.performRegister(request)
	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *registerUserHandlerSuite) TestCreate_InvalidToken() {
	data := fixtures.GetRegisterRequest()
	request := testingHelper.GetRequestBody(data)

	s.registerManager.On("Register", mock.Anything).
		Return(nil, "", errors.New(managers.TokenError)).
		Once()

	recorder := s.performRegister(request)
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
}

func (s *registerUserHandlerSuite) TestCreate() {
	data := fixtures.GetRegisterRequest()
	request := testingHelper.GetRequestBody(data)

	acc := accountFixtures.GetAccount()

	expectedToken := "12345"
	s.registerManager.On("Register", mock.Anything).
		Return(acc, expectedToken, nil).
		Once()

	recorder := s.performRegister(request)
	s.Equal(http.StatusCreated, recorder.Code)
}

func (s *registerUserHandlerSuite) performRegister(body *bytes.Buffer) *httptest.ResponseRecorder {
	return testingHelper.PerformRequest(
		s.handler.Create,
		body,
		http.MethodPost,
		"/register",
	)
}
