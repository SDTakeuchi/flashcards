package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SDTakeuchi/flashcards/adapter/config"
	"github.com/SDTakeuchi/flashcards/adapter/domain_impl/repo"
	"github.com/SDTakeuchi/flashcards/handler"
	"github.com/SDTakeuchi/flashcards/pkg/google/spreadsheet"
	"github.com/SDTakeuchi/flashcards/server"
	"github.com/SDTakeuchi/flashcards/usecase"
)

const defaultAddress = ":8000"

func main() {
	ctx := context.Background()
	sheetService, err := spreadsheet.InitService(
		ctx,
		fmt.Sprintf("./pkg/google/spreadsheet/credential/%s", config.Get().SheetCredential),
		config.Get().TabName,
	)
	if err != nil {
		panic(err)
	}
	cRepo := repo.NewCardRepo(sheetService)
	uuc := usecase.NewUserUsecase()
	cuc := usecase.NewCardUsecase(cRepo)
	uh := handler.NewUserHandler(uuc)
	ch := handler.NewCardHandler(cuc)
	s := server.NewServer(ch, uh)

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	s.Run(address)
}
