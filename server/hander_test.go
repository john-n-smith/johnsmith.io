package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/john-n-smith/johnsmith.io/entry"
	"github.com/stretchr/testify/assert"
)

type mockLoader struct {
	load func(id string) (*entry.Entry, error)
}

func (ml mockLoader) Load(id string) (*entry.Entry, error) {
	return ml.load(id)
}

func TestHandlerEntry(t *testing.T) {
	sampleEntry := &entry.Entry{}
	cases := []struct {
		name       string
		loader     func(id string) (*entry.Entry, error)
		redirector redirector
		renderer   renderer
	}{
		{
			name: "successful",
			loader: func(id string) (*entry.Entry, error) {
				assert.Equal(t, "entry-id", id)
				return sampleEntry, nil
			},
			renderer: func(path string, data interface{}, w http.ResponseWriter) error {
				assert.Equal(t, "./template/entry.html", path)
				assert.Exactly(t, sampleEntry, data)
				return nil
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := server{
				entryLoader: &mockLoader{load: tt.loader},
				redirect:    tt.redirector,
				render:      tt.renderer,
			}

			server.handlerEntry(&httptest.ResponseRecorder{}, &http.Request{URL: &url.URL{Path: "/entry/entry-id"}})
		})
	}
}
