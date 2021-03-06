package ast

import (
	"testing"
)

func TestEnumType(t *testing.T) {
	nodes := map[string]Node{
		`0x7f980b858309 'foo'`: &EnumType{
			Address:  "0x7f980b858309",
			Name:     "foo",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
