package mapdb

import (
	"encoding/json"
	"fmt"
	"os"
	"raspi/boreales/tamanager/sgraph"
)

type DB struct {
	Data map[string]*sgraph.Node
}

func New() *DB {
	return &DB{
		Data: make(map[string]*sgraph.Node),
	}
}
func makeKey(key, ntype string) string {
	return fmt.Sprintf("%s-%s", key, ntype)
}

func (db *DB) Get(key, ntype string) (*sgraph.Node, error) {
	v, ok := db.Data[makeKey(key, ntype)]
	if !ok {
		return v, sgraph.ErrNotFound
	}
	return v, nil
}

func (db *DB) Set(n *sgraph.Node) error {
	db.Data[makeKey(n.Key, n.NodeType)] = n
	return nil
}

func (db *DB) Load(fname string) error {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	return dec.Decode(db)
}

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
