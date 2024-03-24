package server

import (
	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"
	"github.com/SDTakeuchi/go/src/flashcards/handler"
	"github.com/SDTakeuchi/go/src/flashcards/handler/middleware"
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

func (s *Server) Run(address string, tokenIssuer auth.TokenIssuer) {
	engine := gin.Default()

	auth := engine.Group("/backend_api/auth")
	{
		auth.POST("/signup", s.uh.SignUp)
		auth.POST("/login", s.uh.Login)
	}

	api := engine.Group("/backend_api").Use(middleware.CheckAuth(tokenIssuer))
	{
		api.GET("/flashcard", s.ch.GetFlashcard)
		api.GET("/flashcard/remembered", s.ch.GetRememberedFlashcards)
		api.PUT("/flashcard", s.ch.UpdateFlashcard)
	}

	engine.Run(address)
}
