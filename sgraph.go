package sgraph

import "errors"

var ErrNotFound = errors.New("value not found")

type Relation struct {
	Title  string
	KeyTo  string
	TypeTo string
	//Value    []byte
	Property string
}

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

func (g *Graph) GetNode(key, nType string) (*Node, error) {
	return g.DB.Get(key, nType)
}

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

func (g *Graph) Save(fname string) error {
	return g.DB.Save(fname)
}

func (g *Graph) Load(fname string) error {
	return g.DB.Load(fname)
}
