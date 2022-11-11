package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/encoding/dot"
)

type expected struct {
	nodes int
	edges int
}

func TestGraph(t *testing.T) {
	tt := []struct {
		name string
		expected
		graph func(t *testing.T) string
	}{
		{name: "base example", expected: expected{nodes: 4, edges: 3}, graph: simpleGraph},
		{name: "single vertice", expected: expected{nodes: 1, edges: 0}, graph: singleVertice},
		{name: "single edge", expected: expected{nodes: 2, edges: 1}, graph: singleEdge},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			dst := NewGraph()

			if err := dot.Unmarshal([]byte(tc.graph(t)), dst); err != nil {
				fmt.Println(err)
				t.FailNow()
			}

			nodes := dst.Nodes()
			assert.Equalf(t, nodes.Len(), tc.expected.nodes, "length of nodes: %d does not match expected: %d", nodes.Len(), tc.expected.nodes)

			edges := dst.Edges()
			assert.Equalf(t, edges.Len(), tc.expected.edges, "length of edges: %d does not match expected: %d", edges.Len(), tc.expected.edges)
		})
	}
}

func simpleGraph(t *testing.T) string {
	const ug = `
graph {
	a--d [color=black]
	a--b [color=red]
	a--c [color=blue]
}
`
	return ug
}

func singleVertice(t *testing.T) string {
	const ug = `
graph {
	a
}
`
	return ug
}

func singleEdge(t *testing.T) string {
	const ug = `
graph {
	a--b [color=black]
}
`
	return ug
}
