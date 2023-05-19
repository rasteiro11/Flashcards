package main

import (
	"flashcards/core/database"
	"flashcards/core/server"
	"flashcards/models"
	flashcardsHttp "flashcards/src/flashcards/delivery/http"
	flashcardsRepo "flashcards/src/flashcards/repository"
	flashcardsCase "flashcards/src/flashcards/usecase"
	"log"
)

func main() {
	database, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		log.Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

	if err := database.Migrate(models.GetEntities()...); err != nil {
		log.Fatalf("[main] database.Migrate() retunrned error: %+v\n", err)
	}

	server := server.NewServer(server.WithPrefix("/flashcard"))

	db := database.Conn()

	flashcardsRepo := flashcardsRepo.NewRepository(db)

	flashcardsUsecase := flashcardsCase.NewUsecase(
		flashcardsCase.WithRepository(flashcardsRepo),
	)

	flashcardsHttp.NewHandler(server, flashcardsHttp.WithUsecase(flashcardsUsecase))

	server.PrintRouter()

	if err := server.Start(":6969"); err != nil {
		log.Fatalf("[main] server.NewServer() retrurned error: %+v\n", err)
	}
}
