package managers

import (
	"chat/accounts/ports"
	"chat/accounts/ports/mocks"
	"chat/app/accounts/adapters/repositories/fixtures"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type loginManagerSuite struct {
	suite.Suite
	accManager *mocks.AccountManager
	token      *mocks.TokenManager

	manager ports.LoginManager
}

func TestLoginManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(loginManagerSuite))
}

func (s *loginManagerSuite) SetupSuite() {
	s.accManager = &mocks.AccountManager{}
	s.token = &mocks.TokenManager{}
	s.manager = newLoginManager(s.accManager, s.token)
}

func (s *loginManagerSuite) TestLoginAccountNotFound() {
	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Login", mock.Anything).
		Return(nil, errors.New(userNotFound)).
		Once()

	_, token, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(userNotFound, err.Error())
	s.Empty(token)
}

func (s *loginManagerSuite) TestLoginAccountTokenError() {

	e := fixtures.GetAccount()
	acc := e.ToDomain()

	s.accManager.On("Login", mock.Anything).
		Return(acc, nil).
		Once()

	s.token.On("Generate", mock.Anything).
		Return("", errors.New(tokenError)).
		Once()

	_, token, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(tokenError, err.Error())
	s.Empty(token)
}
