package handler

import (
	"net/http"
	"time"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model"
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

type signupRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := model.NewPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	if err := h.userUseCase.SignUp(ctx, req.Name, password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_exp"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_exp"`
	UserID                string    `json:"user_id"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	output, err := h.userUseCase.Login(
		ctx,
		&usecase.LoginInput{
			Name:        req.Name,
			RawPassword: req.Password,
			ClientIP:    c.ClientIP(),
			UserAgent:   c.GetHeader("User-Agent"),
		},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resp := &LoginResponse{
		AccessToken:           output.AccessToken,
		AccessTokenExpiresAt:  output.AccessTokenExpiresAt,
		RefreshToken:          output.RefreshToken,
		RefreshTokenExpiresAt: output.RefreshTokenExpiresAt,
		UserID:                output.UserID.String(),
	}

	c.JSON(http.StatusOK, resp)
}
