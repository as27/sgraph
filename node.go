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

func NewNode(key, nType string, v []byte, rs ...Relation) *Node {
	return &Node{
		Key:       key,
		NodeType:  nType,
		Value:     v,
		Relations: rs,
	}
}

func CreateNode(key, nType string, m Marshaler, rs ...Relation) (*Node, error) {
	b, err := m.Marshal()
	n := NewNode(key, nType, b, rs...)
	return n, err
}

func (n *Node) GetRelations(title string) Relations {
	var rs Relations
	for _, r := range n.Relations {
		if r.Title == title {
			rs = append(rs, r)
		}
	}
	return rs
}

func (n *Node) SetRelation(rs Relation) {
	for _, r := range n.Relations {
		// Don't add when exists
		if reflect.DeepEqual(r, rs) {
			return
		}
	}
	n.Relations = append(n.Relations, rs)
}
