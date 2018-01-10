package sgraph_test

import (
	"bytes"
	"encoding/gob"
	"raspi/boreales/tamanager/sgraph"
	"raspi/boreales/tamanager/sgraph/mapdb"
	"reflect"
	"testing"
)

func TestSgraph(t *testing.T) {
	t1 := MyType{"Alice", 64}
	t2 := MyType{"Bob", 12}
	t3 := MyType{"Charlie", 9}
	graph := sgraph.NewGraph(mapdb.New())
	n1, _ := sgraph.CreateNode("Alice", "MyType", &t1,
		sgraph.Relation{"is_mother_of", "Bob", "MyType", "prop1"},
		sgraph.Relation{"is_mother_of", "Charlie", "MyType", "prop2"},
	)
	n2, _ := sgraph.CreateNode("Bob", "MyType", &t2)
	n3, _ := sgraph.CreateNode("Charlie", "MyType", &t3)
	graph.SetNode(n1)
	graph.SetNode(n2)
	graph.SetNode(n3)

	children, _ := graph.GetConnectedNodes(n1, "is_mother_of")
	expectChildren := []MyType{t2, t3}
	for i, c := range children {
		var n MyType
		n.Unmarshal(c.Value)
		if !reflect.DeepEqual(expectChildren[i], n) {
			t.Error("connected node is not correct")
		}
	}
}

type MyType struct {
	Name string
	Age  int
}

func (mt *MyType) Marshal() ([]byte, error) {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	err := enc.Encode(mt)
	return b.Bytes(), err
}

func (mt *MyType) Unmarshal(b []byte) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	return dec.Decode(mt)
}
