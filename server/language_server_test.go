package server

import (
	"os"
	"testing"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/platform/language"
	"github.com/polyglottis/platform/language/test"
)

var testDB = "test.db"

func TestExists(t *testing.T) {
	os.Remove(testDB)

	db, err := database.Open(testDB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(testDB)

	err = db.Insert("English", &language.English)
	if err != nil {
		t.Fatal(err)
	}

	s := &Server{db: db}

	test.All(s, t)
}
