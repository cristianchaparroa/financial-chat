package rest

import (
	accountInitializer "chat/app/accounts/initializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type server struct {
	router *gin.Engine
}

func NewServer() *server {
	r := gin.Default()
	s := &server{router: r}
	s.setup()
	s.migrations()
	return s
}

func (s *server) setup() {
	s.setupRoutes()
	setupLogger(s)
}

func (s *server) setupRoutes() {
	setupLoginRoutes(s)
}

func (s *server) migrations() {
	accountInitializer.Migrate()
}

func (s *server) Run() {
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)
	s.router.Run(address)
}
