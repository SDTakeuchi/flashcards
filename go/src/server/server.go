package server

import (
	"github.com/SDTakeuchi/go/src/flashcards/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ch *handler.CardHandler
	uh *handler.UserHandler
}

func NewServer(ch *handler.CardHandler, uh *handler.UserHandler) *Server {
	return &Server{
		ch: ch,
		uh: uh,
	}
}

func (s *Server) Run(address string) {
	engine := gin.Default()
	api := engine.Group("/api")
	api.GET("/flashcard", s.ch.GetFlashcard)
	api.PUT("/flashcard", s.ch.UpdateFlashcard)

	user := api.Group("/user")
	{
		user.POST("/sign_up", s.uh.SignUp)
		user.POST("/login", s.uh.Login)
	}
	engine.Run(address)
}
