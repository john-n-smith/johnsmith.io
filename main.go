package main

import (
	"github.com/john-n-smith/johnsmith.io/server"
)

func main() {
	s := server.New()
	s.Serve()
}
