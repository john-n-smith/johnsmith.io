package entry

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/john-n-smith/johnsmith.io/storage"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name        string
		mockStore   storage.Store
		entryID     string
		want        *Entry
		wantErr     bool
		wantErrText string
	}{
		{
			name:      "successful load",
			mockStore: mockStore{},
			entryID:   "my id",
			want:      &Entry{Question: "what is my id?", Answer: "my id"},
		},
		{
			name:        "store error",
			mockStore:   mockStore{err: fmt.Errorf("store error")},
			entryID:     "my id",
			wantErr:     true,
			wantErrText: "store error",
		},
		{
			name: "invalid json from store",
			mockStore: mockStore{fetch: func(id string) ([]byte, error) {
				return []byte("bad json string"), nil
			}},
			wantErr:     true,
			wantErrText: "json unmarshall error: invalid character 'b' looking for beginning of value",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l := loader{store: tt.mockStore}
			e, err := l.Load(tt.entryID)
			if err != nil {
				assert.True(t, tt.wantErr, "we were not expecting an error")
				assert.Equal(t, tt.wantErrText, err.Error())
				return
			}

			assert.False(t, tt.wantErr, "we were expecting an error")
			assert.Equal(t, tt.want, e)
		})
	}
}

type mockStore struct {
	fetch func(id string) ([]byte, error)
	err   error
}

func (ms mockStore) Fetch(id string) ([]byte, error) {
	// if a fetch func has been defined, just call that
	if ms.fetch != nil {
		return ms.fetch(id)
	}

	// otherwise, generate some json
	var jsonB []byte
	if ms.err == nil {
		// return the id as Entry.Answer
		jsonB, _ = json.Marshal(Entry{"what is my id?", id})
	}

	return jsonB, ms.err
}
