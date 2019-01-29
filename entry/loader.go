package entry

import (
	"encoding/json"

	"github.com/john-n-smith/johnsmith.io/config"
	"github.com/john-n-smith/johnsmith.io/storage"
	"github.com/pkg/errors"
)

// Loader holds methods for loading Entry structs
type Loader interface {
	Load(id string) (*Entry, error)
}

type loader struct {
	store storage.Store
}

// Load populates and returns Entry structs
func (l *loader) Load(id string) (*Entry, error) {
	data, err := l.store.Fetch(id)
	if err != nil {
		return nil, err
	}

	e := &Entry{}
	if err := json.Unmarshal(data, e); err != nil {
		return nil, errors.Wrap(err, "json unmarshall error")
	}

	return e, nil
}

// NewLoader returns a configured Loader
func NewLoader(c *config.Configuration) Loader {
	s := storage.New(c)
	return &loader{store: s}
}
