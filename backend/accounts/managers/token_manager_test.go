package managers

import (
	"chat/accounts/ports"
	"chat/accounts/ports/fixtures"
	"github.com/stretchr/testify/suite"
	"testing"
)

type tokeManagerSuite struct {
	suite.Suite
	manager ports.TokenManager
}

func TestTokenManagerSuiteInit(t *testing.T) {
	suite.Run(t, new(tokeManagerSuite))
}

func (s *tokeManagerSuite) SetupSuite() {
	s.manager = newTokenManager()
}

func (s *tokeManagerSuite) TestGenerate() {
	acc := fixtures.GetAccount()
	token, err := s.manager.Generate(acc)
	s.Nil(err)
	s.NotEmpty(token)
}
