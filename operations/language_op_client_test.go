package operations

import (
	"fmt"
	"os"
	"testing"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/platform/language"
)

func TestClientOperationServer(t *testing.T) {
	file := "languages_test.db"
	testAddr := ":1234"

	os.Remove(file)
	db, err := database.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(file)

	op := NewOpServer(db, testAddr)
	if err != nil {
		t.Fatal(err)
	}

	err = op.RegisterAndListen()
	if err != nil {
		t.Fatal(err)
	}

	go op.Accept()

	c, err := NewClient(testAddr)
	if err != nil {
		t.Fatal(err)
	}

	d, err := c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 0 {
		t.Error("Database should be empty at this point")
	}

	err = c.Insert("Test", &language.Language{
		Code:      language.Unknown.Code,
		ISO_693_1: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	d, err = c.Dump()
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 1 {
		t.Fatal("Database should contain exactly one line at this point")
	}
	if d[0] != fmt.Sprintf("%s\ttest\t\t\tTest", language.Unknown.Code) {
		t.Error("Incorrect database dump")
	}
}
