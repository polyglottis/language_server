// Package main contains the language-server executable.
package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/language_server/operations"
	"github.com/polyglottis/language_server/server"
	"github.com/polyglottis/rpc"
)

var dbFile = flag.String("db", "languages.db", "path to sqlite db file")
var tcpAddr = flag.String("tcp", ":14342", "TCP address of language server")
var operationsAddr = flag.String("op-tcp", ":16653", "TCP address of operations RPC server")

func main() {
	flag.Parse()

	abs, err := filepath.Abs(*dbFile)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.Open(abs)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Language server accessing db file %v", abs)

	main := server.New(db, *tcpAddr)
	op := operations.NewOpServer(db, *operationsAddr)
	p := rpc.NewServerPair("Language Server", main, op)

	err = p.RegisterAndListen()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	p.Accept()
}
