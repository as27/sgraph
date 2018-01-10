package sgraph

import "reflect"

type Unmarshaler interface {
	Unmarshal([]byte) error
}

type Marshaler interface {
	Marshal() ([]byte, error)
}

// Node is the basic type for storing data. The struct is stored as  a
// []byte. When writing and reading the data needs to be marshaled or
// unmarshaled.
type Node struct {
	Key       string
	NodeType  string
	Value     []byte
	Relations Relations
}

// NewNode returns a new node
func NewNode(key, nType string, v []byte, rs ...Relation) *Node {
	return &Node{
		Key:       key,
		NodeType:  nType,
		Value:     v,
		Relations: rs,
	}
}

// CreateNode creates a node, by using the marshaler interface.
func CreateNode(key, nType string, m Marshaler, rs ...Relation) (*Node, error) {
	b, err := m.Marshal()
	n := NewNode(key, nType, b, rs...)
	return n, err
}

// GetRelations returns the relations for a given relation type
func (n *Node) GetRelations(title string) Relations {
	var rs Relations
	for _, r := range n.Relations {
		if r.Title == title {
			rs = append(rs, r)
		}
	}
	return rs
}

// SetRelation adds a new relation to the node.
func (n *Node) SetRelation(rs Relation) {
	for _, r := range n.Relations {
		// Don't add when exists
		if reflect.DeepEqual(r, rs) {
			return
		}
	}
	n.Relations = append(n.Relations, rs)
}
