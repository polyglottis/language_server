package operations

import (
	"log"

	"github.com/polyglottis/language_server/database"
	"github.com/polyglottis/platform/language"
	"github.com/polyglottis/rpc"
)

type OpRpcServer struct {
	db *database.DB
}

// NewOpServer creates a new rpc server for maintenance operations on the language server.
func NewOpServer(db *database.DB, addr string) *rpc.Server {
	return rpc.NewServer("OpRpcServer", &OpRpcServer{db}, addr)
}

type InsertArgs struct {
	Comment   string
	Code      string
	ISO_639_1 string
	ISO_639_3 string
	WikiData  string
}

func (s *OpRpcServer) Insert(args InsertArgs, nothing *bool) error {
	log.Printf("About to insert %+v", args)
	err := s.db.Insert(args.Comment, &language.Language{
		Code:      language.Code(args.Code),
		ISO_639_1: args.ISO_639_1,
		ISO_639_3: args.ISO_639_3,
		WikiData:  args.WikiData,
	})
	log.Printf("Insertion returned %v", err)
	return err
}

type Dump []string

func (s *OpRpcServer) Dump(nothing bool, d *Dump) error {
	lines, err := s.db.Dump()
	if err != nil {
		return err
	}
	*d = lines
	return nil
}
