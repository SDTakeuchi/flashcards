package handler

import (
	"net/http"

	"github.com/SDTakeuchi/go/src/flashcards/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUsecase
}

func NewUserHandler(userUseCase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
