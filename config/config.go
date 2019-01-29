package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

const prefix = "JSIO_"

var (
	conf Configuration
	once sync.Once
)

// Configuration holds the application's configuration
type Configuration struct {
	Storage string
}

// Config returns the application's configuration. Env vars are parsed only once on initial invocation
func Config() *Configuration {
	var err error
	once.Do(func() {
		err = parse(&conf)
	})
	if err != nil {
		panic(err)
	}

	return &conf
}

func parse(c *Configuration) (err error) {
	s := reflect.ValueOf(c).Elem()
	t := s.Type()

	for i := 0; i < s.NumField(); i++ {
		fn := t.Field(i).Name
		en := fieldNameToEnvName(fn)

		if val, ok := os.LookupEnv(en); ok {
			f := s.Field(i)
			f.Set(reflect.ValueOf(val))
			continue
		}

		err = fmt.Errorf("env var not set: %s", en)
	}

	return
}

// fieldNameToEnvName takes a camelcase field name and returns the equivalent env var name
func fieldNameToEnvName(fn string) string {
	s := prefix

	runes := []rune(fn)
	for i := 0; i < len(runes); i++ {
		// the current rune is upper and either the prev or next is lower, insert underscore
		if i > 0 && unicode.IsUpper(runes[i]) && (unicode.IsLower(runes[i-1]) || (i < len(runes)-1 && unicode.IsLower(runes[i+1]))) {
			s = s + "_"
		}

		s = s + string(runes[i])
	}

	return strings.ToUpper(s)
}
