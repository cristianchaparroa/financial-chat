package managers

import (
	"chat/accounts"
	"chat/accounts/ports"
	"chat/accounts/ports/fixtures"
	"chat/accounts/ports/mocks"
	entityFixtures "chat/app/accounts/adapters/repositories/fixtures"
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

func (s *accountManagerSuite) TestLoginUserNotAuthenticated() {
	entity := entityFixtures.GetAccount()
	acc := entity.ToDomain()

	s.reader.On("FindByEmail", mock.Anything).
		Return(acc, nil).
		Once()

	acc, err := s.manager.Login(acc)
	s.NotNil(err)
	s.Equal(userNotAuthenticated, err.Error())
}

func (s *accountManagerSuite) TestLoginUserAuthenticated() {
	entity := entityFixtures.GetAccount()
	entity.Password = "$2a$04$OM9mX7jdK34ey6Tu8gfx9eGbXVECy/x9zsHVNehkIb3DUvldOcAOq"
	accountRequest := entity.ToDomain()

	password := "De76cBmu6AhGJNWzSXUXtaJkRhTj7EQ9jff46A7sXXg9Q"
	acc := &accounts.Account{Password: password}

	s.reader.On("FindByEmail", mock.Anything).
		Return(accountRequest, nil).
		Once()

	acc, err := s.manager.Login(acc)
	s.Nil(err)
	s.NotNil(acc)
}

func (s *accountManagerSuite) TestCreate() {
	entity := entityFixtures.GetAccount()
	acc := entity.ToDomain()

	s.writer.On("Create", mock.Anything).Return(acc, nil).Once()

	result, err := s.manager.Create(acc)
	s.Nil(err)
	s.Equal(acc, result)
}

func (s *accountManagerSuite) TestGetByID() {
	entity := entityFixtures.GetAccount()
	acc := entity.ToDomain()

	s.reader.On("FindByID", mock.Anything).
		Return(acc).
		Once()

	result := s.manager.GetByID("12345")
	s.NotNil(result)
}
