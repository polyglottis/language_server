// Package server defines the Polyglottis Language Server.
package server

import (
	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/platform/language"
	languageRpc "github.com/polyglottis/platform/language/rpc"
	"github.com/polyglottis/rpc"
)

type Server struct {
	db *database.DB
}

func NewServerDB(db *database.DB) *Server {
	return &Server{db: db}
}

func NewServer(dbFile string) (*Server, error) {
	db, err := database.Open(dbFile)
	if err != nil {
		return nil, err
	}
	return NewServerDB(db), nil
}

// New creates the rpc language server, as required by polyglottis/language/rpc
func New(s *Server, serverAddr string) *rpc.Server {
	return languageRpc.NewLanguageServer(s, serverAddr)
}

func (s *Server) GetCode(code string) (language.Code, error) {
	exists, err := s.db.CodeExists(code)
	if err != nil {
		return "", err
	}
	if exists {
		return language.Code(code), nil
	} else {
		return language.Unknown.Code, language.CodeNotFound
	}
}

func (s *Server) List() ([]language.Code, error) {
	return s.db.CodeList()
}
