package repositories

import (
	"chat/app/chat/adapters/repositories/fixtures"
	"chat/app/dataproviders/sql"
	"chat/chat/ports"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userWriterSuite struct {
	suite.Suite
	writer ports.UserWriter
}

func TestUserWriterInit(t *testing.T) {
	suite.Run(t, new(userWriterSuite))
}

func (s *userWriterSuite) SetupTest() {
	conn := sql.NewMockConnection()
	s.writer = newUserWriter(conn)
}

func (s *userWriterSuite) TestUserWriter_Create() {
	u := fixtures.GetUser()
	user := s.writer.Create(u.ToDomain())
	s.Equal(u.ToDomain(), user)
}
