// Package mapdb implemnents the sgraph.DB interface. It is a
// simple implementation by just using a map to store the data
// in memory.
package mapdb

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/as27/sgraph"
)

// DB is the type using
type DB struct {
	data map[string]*sgraph.Node
}

// New creates a new instance of the mapdb
func New() *DB {
	return &DB{
		data: make(map[string]*sgraph.Node),
	}
}

func makeKey(key, ntype string) string {
	return fmt.Sprintf("%s-%s", key, ntype)
}

// Get returns a Node to a given key
func (db *DB) Get(key, ntype string) (*sgraph.Node, error) {
	v, ok := db.data[makeKey(key, ntype)]
	if !ok {
		return v, sgraph.ErrNotFound
	}
	return v, nil
}

// Set stores a node inside the db
func (db *DB) Set(n *sgraph.Node) error {
	db.data[makeKey(n.Key, n.NodeType)] = n
	return nil
}

// Load the data from a file
func (db *DB) Load(fname string) error {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	return dec.Decode(db)
}

// Save the data to a file
func (db *DB) Save(fname string) error {
	f, err := os.OpenFile(fname, os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(db)
	return nil
}
