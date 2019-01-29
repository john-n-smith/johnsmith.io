package config

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldNameToEnvName(t *testing.T) {
	cases := []struct {
		name      string
		fieldName string
		want      string
	}{
		{
			name:      "standard field name",
			fieldName: "DataStore",
			want:      "JSIO_DATA_STORE",
		},
		{
			name:      "accronym prefix and suffix",
			fieldName: "BBQSauceJSON",
			want:      "JSIO_BBQ_SAUCE_JSON",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// reset once to ensure env vars are parsed again
			once = sync.Once{}
			en := fieldNameToEnvName(tt.fieldName)
			assert.Equal(t, tt.want, en)
		})
	}
}

func TestConfig(t *testing.T) {
	cases := []struct {
		name         string
		env          map[string]string
		want         *Configuration
		wantPanic    bool
		wantPanicMsg string
	}{
		{
			name: "success all required envs",
			env: map[string]string{
				"JSIO_STORAGE": "i am storage",
			},
			want: &Configuration{
				Storage: "i am storage",
			},
		},
		{
			name:         "failure missing env",
			env:          map[string]string{},
			want:         &Configuration{},
			wantPanic:    true,
			wantPanicMsg: "env var not set: JSIO_STORAGE",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					assert.True(t, tt.wantPanic, "we were not expecting a panic")
					assert.Equal(t, tt.wantPanicMsg, r.(error).Error())
				} else {
					assert.False(t, tt.wantPanic, "we were expecting a panic")
				}

				for k := range tt.env {
					os.Unsetenv(k)
				}
			}()

			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			// reset once to ensure env vars are parsed again
			once = sync.Once{}

			conf := Config()

			assert.Equal(t, tt.want, conf)
		})
	}
}

func TestConfigPanic(t *testing.T) {
	cases := []struct {
		name        string
		env         map[string]string
		wantErrText string
	}{
		{
			name:        "no envs",
			wantErrText: "env var not set: JSIO_STORAGE",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						assert.Equal(t, tt.wantErrText, err.Error())
					} else {
						assert.Fail(t, "expecting error from recovery")
					}
				}
			}()
			// define sync here to ensure we load config again
			once = sync.Once{}
			Config()
		})
	}
}
