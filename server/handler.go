package server

import (
	"fmt"
	"net/http"
)

func (s *server) handlerEntry(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/entry/"):]

	entry, err := s.entryLoader.Load(id)
	if err != nil {
		s.redirect(w, r, "/four-oh-four", http.StatusNotFound)
		return
	}

	if err := s.render("./template/entry.html", entry, w); err != nil {
		fmt.Println(err)
	}
}

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sorry, couldn't find that page")
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I like %s", r.URL.Path[1:])
}
