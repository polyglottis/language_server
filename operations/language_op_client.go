// Package operations contains an rpc client-server pair for maintenance operations on the language server.
package operations

import (
	"log"
	"net/rpc"

	"github.com/polyglottis/platform/language"
)

type Client struct {
	c *rpc.Client
}

// NewClient creates an rpc client for maintenance operations on the language server.
func NewClient(addr string) (*Client, error) {
	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{c: c}, nil
}

// Insert inserts a new language into the language database, with the desired comment.
func (c *Client) Insert(comment string, language *language.Language) error {
	log.Printf("Inserting %v, %+v", comment, language)
	return c.c.Call("OpRpcServer.Insert", InsertArgs{
		Comment:   comment,
		Code:      string(language.Code),
		ISO_693_1: language.ISO_693_1,
		ISO_693_3: language.ISO_693_3,
		ISO_693_6: language.ISO_693_6,
	}, nil)
}

// Dump returns a dump of the whole language database.
func (c *Client) Dump() ([]string, error) {
	var d Dump
	err := c.c.Call("OpRpcServer.Dump", false, &d)
	if err != nil {
		return nil, err
	}
	return []string(d), nil
}
