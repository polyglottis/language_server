// Package main contains the language-server executable.
package main

import (
	"log"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/language_server/operations"
	"github.com/polyglottis/language_server/server"
	"github.com/polyglottis/platform/config"
	"github.com/polyglottis/rpc"
)

func main() {

	c := config.Get()

	db, err := database.Open(c.LanguageDB)
	if err != nil {
		log.Fatalln(err)
	}

	main := server.New(server.NewServerDB(db), c.Language)
	op := operations.NewOpServer(db, c.LanguageOp)
	p := rpc.NewServerPair("Language Server", main, op)

	err = p.RegisterAndListen()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	p.Accept()
}
