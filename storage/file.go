package storage

import (
	"io/ioutil"
	"log"
)

// storeFile is an implementation of Store for the filesystem
type storeFile struct {
}

// Fetch retrieves data from the filesystem, identified by id
func (s storeFile) Fetch(id string) ([]byte, error) {
	filepath := "./db/" + id + ".json"
	log.Println("attempting to load from path", filepath)
	return ioutil.ReadFile(filepath)
}
