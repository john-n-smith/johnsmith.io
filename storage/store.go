package storage

import (
	"fmt"

	"github.com/john-n-smith/johnsmith.io/config"
)

// Store is somewhere to retrieve Entry data from
type Store interface {
	Fetch(id string) ([]byte, error)
}

// New returns a Store interface
func New(c *config.Configuration) Store {
	var s Store

	switch c.Storage {
	case "FILE":
		s = &storeFile{}
	case "DYNAMO_DB":
		s = &storeDynamo{}
	default:
		panic(fmt.Errorf("Unknown store: %s", c.Storage))
	}

	return s
}
