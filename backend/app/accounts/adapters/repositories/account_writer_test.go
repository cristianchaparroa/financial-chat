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
	user := s.writer.Create(u.ToDomain())
	s.Equal(u.ToDomain(), user)
}
