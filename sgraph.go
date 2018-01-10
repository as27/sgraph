package sgraph

import "errors"

var ErrNotFound = errors.New("value not found")

// Relation is used to connect a node to another. The connection
// is stored at a node and points into the direction of the other
// node.
type Relation struct {
	Title  string
	KeyTo  string
	TypeTo string
	//Value    []byte
	Property string
}

// NewRelation creates a new relation, which can be added to a node
func NewRelation(title, key, nType, prop string) Relation {
	//value, err := m.Marshal()
	return Relation{title, key, nType, prop}
}

type Relations []Relation

// DB is the interface for storing and getting the Date
type DB interface {
	Get(key, ntype string) (*Node, error)
	Set(n *Node) error
	Load(fname string) error
	Save(fname string) error
}

// Graph is the simple db
type Graph struct {
	DB
}

// NewGraph takes a object which implements the DB interface
func NewGraph(db DB) *Graph {
	return &Graph{db}
}

// CreateNode creates a node and adds it to the graph
func (g *Graph) CreateNode(key, nType string, m Marshaler, rs ...Relation) error {
	n, err := CreateNode(key, nType, m, rs...)
	if err != nil {
		return err
	}
	return g.DB.Set(n)
}

// SetNode sets the pointer to a node.
func (g *Graph) SetNode(n *Node) error {
	return g.DB.Set(n)
}

// GetNode returns a pointer to a node
func (g *Graph) GetNode(key, nType string) (*Node, error) {
	return g.DB.Get(key, nType)
}

// GetConnectedNodes returns a slice of the connected, which have the given
// relation.
func (g *Graph) GetConnectedNodes(n *Node, relation string) ([]*Node, error) {
	var ns []*Node
	rs := n.GetRelations(relation)
	for _, r := range rs {
		cn, err := g.GetNode(r.KeyTo, r.TypeTo)
		if err != nil {
			return ns, err
		}
		ns = append(ns, cn)
	}
	return ns, nil
}

// Save stores the db into a file
func (g *Graph) Save(fname string) error {
	return g.DB.Save(fname)
}

// Load loads the db from a file
func (g *Graph) Load(fname string) error {
	return g.DB.Load(fname)
}
