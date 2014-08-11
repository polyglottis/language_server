// Package language-dump contains the language-dump executable.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/polyglottis/language_server/operations"
)

var operationsAddr = flag.String("op-tcp", ":16653", "TCP address of operations RPC server")

func main() {
	flag.Parse()

	c, err := operations.NewClient(*operationsAddr)
	if err != nil {
		log.Fatalln(err)
	}

	lines, err := c.Dump()
	if err != nil {
		log.Fatalln(err)
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
