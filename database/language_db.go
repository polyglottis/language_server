// Package database defines the language database.
package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3" // driver import

	"github.com/polyglottis/platform/database"
	"github.com/polyglottis/platform/language"
)

type DB struct {
	db *database.DB
}

func Open(file string) (*DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	langDB, err := database.Create(db, database.Schema{{
		Name: "languages",
		Columns: database.Columns{{
			Field:      "code",
			Type:       "text",
			Constraint: "primary key not null",
		}, {
			Field: "iso_639_1",
			Type:  "text",
		}, {
			Field: "iso_639_3",
			Type:  "text",
		}, {
			Field: "wikiData",
			Type:  "text",
		}, {
			Field: "comment",
			Type:  "text",
		}},
	}})
	if err != nil {
		return nil, err
	}

	return &DB{
		db: langDB,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) CodeExists(code string) (bool, error) {
	return db.db.QueryNonZero("select count(1) from languages where code=?", code)
}

func (db *DB) Insert(comment string, language *language.Language) error {
	if language == nil {
		return fmt.Errorf("Language should not be nil")
	}
	_, err := db.db.Exec("insert into languages values (?,?,?,?,?)", string(language.Code), language.ISO_639_1, language.ISO_639_3, language.WikiData, comment)
	return err
}

func (db *DB) Dump() ([]string, error) {
	s := make([]string, 0)
	rows, err := db.db.Query("select * from languages")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		vals := make([]interface{}, 5)
		for i := range vals {
			vals[i] = new(string)
		}
		if err := rows.Scan(vals...); err != nil {
			return nil, err
		}
		str := make([]string, 5)
		for i, v := range vals {
			str[i] = *v.(*string)
		}
		s = append(s, strings.Join(str, "\t"))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return s, nil
}

func (db *DB) CodeList() ([]language.Code, error) {
	list := make([]language.Code, 0)
	rows, err := db.db.Query("select code from languages")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			return nil, err
		}
		list = append(list, language.Code(code))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
