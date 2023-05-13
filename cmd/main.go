package main

import (
	"flashcards/src/flashcards/database"
	"log"
)

func main() {
	_, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		log.Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

}
