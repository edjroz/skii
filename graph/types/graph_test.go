package types

import (
	"testing"

	"gonum.org/v1/gonum/graph/encoding/dot"
)

const ug = `
graph {
	a--d [color=black]
	a--b [color=red]
	a--c [color=blue]
}
`

func TestGraph(t *testing.T) {
	dst := NewGraph()
	if err := dot.Unmarshal([]byte(ug), dst); err != nil {
		t.FailNow()
	}
	nodes := dst.Nodes()
	if nodes.Len() != 4 {
		t.FailNow()
	}
	edges := dst.Edges()
	if edges.Len() != 3 {
		t.FailNow()
	}
}
