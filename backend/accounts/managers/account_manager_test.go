package managers

import (
	"chat/accounts/ports"
	"chat/accounts/ports/fixtures"
	"chat/accounts/ports/mocks"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountManagerSuite struct {
	suite.Suite
	writer  *mocks.AccountWriter
	reader  *mocks.AccountReader
	manager ports.AccountManager
}

func TestAccountManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(accountManagerSuite))
}

func (s *accountManagerSuite) SetupSuite() {
	s.reader = &mocks.AccountReader{}
	s.writer = &mocks.AccountWriter{}

	s.manager = newAccountManager(s.writer, s.reader)
}

func (s *accountManagerSuite) TestLoginUserNotfound() {
	acc := fixtures.GetAccount()
	s.reader.On("FindByEmail", mock.Anything).
		Return(nil, errors.New(userNotFound)).
		Once()

	acc, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(userNotFound, err.Error())

}
