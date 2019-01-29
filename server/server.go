package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/john-n-smith/johnsmith.io/config"
	"github.com/john-n-smith/johnsmith.io/entry"
)

type redirector func(w http.ResponseWriter, r *http.Request, url string, code int)

type renderer func(path string, data interface{}, w http.ResponseWriter) error

// Server is an http server
type server struct {
	logger      *log.Logger
	mux         *http.ServeMux
	config      *config.Configuration
	entryLoader entry.Loader
	redirect    redirector
	render      renderer
}

func (s *server) routes() {
	s.mux.HandleFunc("/entry/", s.handlerEntry)
	s.mux.HandleFunc("/four-oh-four", handlerNotFound)
	s.mux.HandleFunc("/", handlerRoot)
	fmt.Println("roots added")
}

// New returns a reference to a new server
func New() *server {
	conf := config.Config()

	return &server{
		&log.Logger{},
		http.NewServeMux(),
		conf,
		entry.NewLoader(conf),
		func(w http.ResponseWriter, r *http.Request, url string, code int) {
			http.Redirect(w, r, url, code)
		},
		func(path string, data interface{}, w http.ResponseWriter) error {
			t, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			return t.Execute(w, data)
		},
	}
}

// Serve adds routes to Server and starts it listening
func (s *server) Serve() {
	fmt.Println("starting server")
	s.routes()
	s.logger.Fatal(http.ListenAndServe(":8081", s.mux))
}
