package operations

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestEncoding(t *testing.T) {
	args := InsertArgs{
		Comment: "test",
		Code:    "testCode",
	}

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	err := enc.Encode(args)
	if err != nil {
		t.Fatal(err)
	}

	var a InsertArgs
	err = dec.Decode(&a)
	if err != nil {
		t.Fatal(err)
	}

	if a.Comment != args.Comment {
		t.Fatal("WTF")
	}
	if a.Code != args.Code {
		t.Fatal("WTF2")
	}
}
