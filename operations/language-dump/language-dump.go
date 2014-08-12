// Package language-dump contains the language-dump executable.
package main

import (
	"fmt"
	"log"

	"github.com/polyglottis/language_server/operations"
	"github.com/polyglottis/platform/config"
)

func main() {

	conf := config.Get()

	c, err := operations.NewClient(conf.LanguageOp)
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
