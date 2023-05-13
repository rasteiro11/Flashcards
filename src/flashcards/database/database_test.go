package database_test

import (
	"flashcards/src/flashcards/database"
	"log"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	database, err := database.NewDatabase(database.GetMysqlEngineBuilder)
	if err != nil {
		log.Fatalf("[main] database.NewDatabase() retunrned error: %+v\n", err)
	}

	if err != nil {
		t.Fatalf("database.NewDatabase retunrned error: %+v\n", err)
	}

	if database == nil {
		t.Errorf("database.NewDatabase() must return something")
	}
}
