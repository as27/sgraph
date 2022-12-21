package mapdb

import (
	"reflect"
	"testing"

	"github.com/as27/sgraph"
)

func TestDBImplementation(t *testing.T) {
	var _ sgraph.DB = New()

}

func TestDB_Get(t *testing.T) {
	testDB := createDB()
	type args struct {
		key   string
		ntype string
	}
	tests := []struct {
		name    string
		db      *DB
		args    args
		want    *sgraph.Node
		wantErr bool
	}{
		{
			"one",
			testDB,
			args{"one", "number"},
			testDB.data[makeKey("one", "number")],
			false,
		},
		{
			"value not found",
			testDB,
			args{"ten", "number"},
			testDB.data[makeKey("ten", "number")],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.Get(tt.args.key, tt.args.ntype)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createNode(key, nType, val string) *sgraph.Node {
	return &sgraph.Node{
		Key:      key,
		NodeType: nType,
		Value:    []byte(val),
	}
}

func createDB() *DB {
	db := New()
	db.Set(createNode("one", "number", "one"))
	db.Set(createNode("two", "number", "two"))
	return db
}

func TestDB_Set(t *testing.T) {
	type args struct {
		key      string
		nodeType string
		value    string
	}
	tests := []struct {
		name    string
		db      *DB
		args    args
		wantErr bool
	}{
		{
			"one",
			createDB(),
			args{"one", "number", "1"},

			false,
		},
		{
			"ten",
			createDB(),
			args{"ten", "number", "10"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			n := createNode(tt.args.key, tt.args.nodeType, tt.args.value)
			err := tt.db.Set(n)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.db.data[makeKey(tt.args.key, tt.args.nodeType)] != n {
				t.Error("Set() is not setting the correct pointer")
			}
		})
	}
}
