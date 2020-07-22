package repositories

import (
	"chat/app/chat/adapters/repositories/fixtures"
	"chat/app/dataproviders/sql"
	"chat/chat/ports"
	"fmt"
	uuid "github.com/satori/go.uuid"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	selectByID = "SELECT * FROM users WHERE ID ="
)

type userReaderSuite struct {
	suite.Suite
	reader ports.UserReader
}

func TestUserReaderInit(t *testing.T) {
	suite.Run(t, new(userReaderSuite))
}

func (s *userReaderSuite) SetupSuite() {
	conn := sql.NewMockConnection()
	s.reader = newUserReader(conn)
}

func (s *userReaderSuite) TestUserReader_FinUserByID() {

	id := uuid.NewV1().String()
	sql := fmt.Sprintf("%s?=%s", selectByID, id)
	userExpected := fixtures.GetUserEntity(id)

	mocket.Catcher.Logging = true
	mocket.Catcher.Reset().NewMock().
		WithQuery(sql).
		WithReply(userExpected)

	user := s.reader.FinUserByID(id)
	s.NotNil(user)
	//s.Equal(id, user.ID)
}
