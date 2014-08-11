package database

import (
	"os"
	"testing"

	"github.com/polyglottis/platform/language"
)

var testDB = "test.db"

func TestExists(t *testing.T) {
	os.Remove(testDB)

	db, err := Open(testDB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(testDB)

	err = db.Insert("English", &language.English)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("English again", &language.Language{
		Code: "en",
	})
	if err == nil {
		t.Errorf("Duplicate language codes should trigger an error: 'en'")
	}

	exists, err := db.CodeExists("en")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Error("Language should exist: 'en'")
	}

	invalid := "invalid language"
	exists, err = db.CodeExists(invalid)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Errorf("Language should not exist: '%s'", invalid)
	}
}
