package storage

// storeDynamo is an implementation of Store for Dynamo DB
type storeDynamo struct {
}

// Fetch retrieves data from Dynamo DB, identified by id
func (l storeDynamo) Fetch(id string) ([]byte, error) {
	// @todo
	return nil, nil
}
