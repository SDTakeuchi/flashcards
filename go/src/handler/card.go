package handler

import (
	"net/http"

	"github.com/SDTakeuchi/flashcards/domain/model"
	"github.com/SDTakeuchi/flashcards/usecase"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	cardUseCase usecase.CardUsecase
}

func NewCardHandler(cardUseCase usecase.CardUsecase) *CardHandler {
	return &CardHandler{
		cardUseCase: cardUseCase,
	}
}

func (h *CardHandler) GetFlashcard(c *gin.Context) {
	ctx := c.Request.Context()
	card, err := h.cardUseCase.GetOldest(ctx)
	if err != nil {
		respError(c, http.StatusInternalServerError, err.Error())
		return
	}
	respSuccess(c, convertCard(card))
}

type UpdateFlashcardRequest struct {
	Word        string `json:"word"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
}

func (h *CardHandler) UpdateFlashcard(c *gin.Context) {
	ctx := c.Request.Context()
	var req UpdateFlashcardRequest
	if err := c.BindJSON(&req); err != nil {
		respError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.cardUseCase.Update(
		ctx,
		req.Word,
		req.Description,
		model.CardStatusFromUint8(req.Status),
	); err != nil {
		respError(c, http.StatusInternalServerError, err.Error())
		return
	}
	respSuccess(c, "update success")
}
