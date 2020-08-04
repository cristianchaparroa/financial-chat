package repositories

import (
	"chat/accounts/ports"
	"chat/app/accounts/adapters/repositories/fixtures"
	"chat/app/dataproviders/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountWriterSuite struct {
	suite.Suite
	writer ports.AccountWriter
}

func TestUserWriterInit(t *testing.T) {
	suite.Run(t, new(accountWriterSuite))
}

func (s *accountWriterSuite) SetupTest() {
	conn := sql.NewMockConnection()
	s.writer = newAccountWriter(conn)
}

func (s *accountWriterSuite) TestAccountWriter_Create() {
	u := fixtures.GetAccount()
	userExpected := u.ToDomain()

	userCreated, err := s.writer.Create(userExpected)
	s.Nil(err)

	userExpected = userExpected.SetHashPassword(userExpected.Password)

	s.Equal(userExpected.Email, userCreated.Email)
	s.Equal(userExpected.Name, userCreated.Name)
	s.Equal(userExpected.CreatedAt, userCreated.CreatedAt)
	s.Equal(userExpected.UpdatedAt, userCreated.UpdatedAt)
}
