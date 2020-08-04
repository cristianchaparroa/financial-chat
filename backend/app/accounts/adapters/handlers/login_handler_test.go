package handlers

import (
	"bytes"
	"chat/accounts/managers"
	accountFixtures "chat/accounts/ports/fixtures"
	"chat/accounts/ports/mocks"
	"chat/app/accounts/adapters/handlers/fixtures"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type loginHandlerSuite struct {
	suite.Suite
	loginManager *mocks.LoginManager
	handler      *LoginHandler
}

func TestLoginHandlerSuiteInit(t *testing.T) {
	suite.Run(t, new(loginHandlerSuite))
}

func (s *loginHandlerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	s.loginManager = &mocks.LoginManager{}
	s.handler = newLoginHandler(s.loginManager)
}

func (s *loginHandlerSuite) TestEmptyLoginRequest() {
	data := fixtures.GetEmptyLoginRequest()
	request := fixtures.GetLoginRequestBody(data)
	recorder := s.performLogin(request)
	s.Equal(http.StatusBadRequest, recorder.Code)
}
func (s *loginHandlerSuite) TestLoginRequest_invalidToken() {
	data := fixtures.GetLoginRequest()
	request := fixtures.GetLoginRequestBody(data)

	acc := accountFixtures.GetAccount()

	s.loginManager.On("Login", mock.Anything).
		Return(acc, "", errors.New(managers.TokenError)).
		Once()

	recorder := s.performLogin(request)
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)
}

func (s *loginHandlerSuite) TestLoginRequest() {
	data := fixtures.GetLoginRequest()
	request := fixtures.GetLoginRequestBody(data)

	acc := accountFixtures.GetAccount()

	expectedToken := "12345"
	s.loginManager.On("Login", mock.Anything).
		Return(acc, expectedToken, nil).
		Once()

	recorder := s.performLogin(request)
	s.Equal(http.StatusOK, recorder.Code)
}

func (s *loginHandlerSuite) performLogin(body *bytes.Buffer) *httptest.ResponseRecorder {
	return s.performRequest(
		s.handler.Login,
		body,
		http.MethodPost,
		"/v1/users/register",
	)
}

func (s *loginHandlerSuite) performRequest(handler gin.HandlerFunc, body *bytes.Buffer, method, path string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(recorder)
	r.Handle(method, path, handler)
	c.Request, _ = http.NewRequest(http.MethodPost, path, body)
	r.ServeHTTP(recorder, c.Request)
	return recorder
}
