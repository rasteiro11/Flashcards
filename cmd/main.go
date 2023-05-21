package main

import (
	"flashcards/core/database"
	"flashcards/core/server"
	middlewares "flashcards/middleware"
	"flashcards/models"
	flashcardsHttp "flashcards/src/flashcards/delivery/http"
	flashcardsRepo "flashcards/src/flashcards/repository"
	flashcardsCase "flashcards/src/flashcards/usecase"
	usersHttp "flashcards/src/user/delivery/http"
	usersRepo "flashcards/src/user/repository"
	usersCase "flashcards/src/user/usecase"
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
	server.Use("/user", middlewares.ValidateUserMiddleware())

	db := database.Conn()

	cardsRepo := flashcardsRepo.NewRepository(db)
	usersRepo := usersRepo.NewRepository(db)

	flashcardsUsecase := flashcardsCase.NewUsecase(
		flashcardsCase.WithRepository(cardsRepo),
	)
	usersUsecase := usersCase.NewUsecase(
		usersCase.WithRepository(usersRepo),
	)

	flashcardsHttp.NewHandler(server, flashcardsHttp.WithUsecase(flashcardsUsecase))
	usersHttp.NewHandler(server, usersHttp.WithUsecase(usersUsecase))

	server.PrintRouter()

	if err := server.Start(":6969"); err != nil {
		log.Fatalf("[main] server.NewServer() retrurned error: %+v\n", err)
	}
}
