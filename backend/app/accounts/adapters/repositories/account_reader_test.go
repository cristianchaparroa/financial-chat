package repositories

import (
	"chat/accounts/ports"
	"chat/app/accounts/adapters/repositories/fixtures"
	"chat/app/dataproviders/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	selectByID    = "SELECT * FROM accounts WHERE ID ="
	selectByEmail = "SELECT * FROM accounts WHERE email ="
)

type accountReaderSuite struct {
	suite.Suite
	reader ports.AccountReader
}

func TestAccountReaderInit(t *testing.T) {
	suite.Run(t, new(accountReaderSuite))
}

func (s *accountReaderSuite) SetupSuite() {
	conn := sql.NewMockConnection()
	s.reader = newAccountReader(conn)
}

func (s *accountReaderSuite) TestAccountReader_FindByID() {

	id := uuid.NewV1().String()
	sql := fmt.Sprintf("%s?=%s", selectByID, id)
	userExpected := fixtures.GetAccountEntity(id)

	mocket.Catcher.Logging = true
	mocket.Catcher.Reset().NewMock().
		WithQuery(sql).
		WithReply(userExpected)

	user := s.reader.FindByID(id)
	s.NotNil(user)
}

func (s *accountReaderSuite) TestAccountReader_FindByEmail() {
	email := "test@gmail.com"
	id := uuid.NewV1().String()

	sql := fmt.Sprintf("%s?=%s", selectByEmail, email)

	accExpected := fixtures.GetAccountEntity(id)

	mocket.Catcher.Logging = true
	mocket.Catcher.Reset().NewMock().
		WithQuery(sql).
		WithReply(accExpected)

	acc := s.reader.FindByEmail(email)
	s.NotNil(acc)
}
