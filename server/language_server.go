// Package server defines the Polyglottis Language Server.
package server

import (
	"errors"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/platform/language"
	localeRpc "github.com/polyglottis/platform/language/rpc"
	"github.com/polyglottis/rpc"
)

var CodeNotFound = errors.New("Language code not found")

type Server struct {
	db *database.DB
}

// New creates the rpc language server, as required by polyglottis/locale/rpc
func New(db *database.DB, serverAddr string) *rpc.Server {
	s := &Server{
		db: db,
	}

	return localeRpc.NewLanguageServer(s, serverAddr)
}

func (s *Server) GetCode(code string) (language.Code, error) {
	exists, err := s.db.CodeExists(code)
	if err != nil {
		return "", err
	}
	if exists {
		return language.Code(code), nil
	} else {
		return language.Unknown.Code, CodeNotFound
	}
}
