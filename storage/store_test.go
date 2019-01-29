package storage

import (
	"testing"

	"github.com/john-n-smith/johnsmith.io/config"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	cases := []struct {
		name   string
		config *config.Configuration
		want   Store
	}{
		{
			name:   "file store",
			config: &config.Configuration{Storage: "FILE"},
			want:   &storeFile{},
		},
		{
			name:   "dynamo db store",
			config: &config.Configuration{Storage: "DYNAMO_DB"},
			want:   &storeDynamo{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.config)
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestNewStorePanic(t *testing.T) {
	cases := []struct {
		name        string
		config      *config.Configuration
		wantErrText string
	}{
		{
			name:   "invalid store",
			config: &config.Configuration{Storage: "INVALID"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panicsf(t, func() { New(tt.config) }, "Unknown store: INVALID")
		})
	}
}
