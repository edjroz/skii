package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type expected struct {
	err   bool
	nodes int
	edges int
}

func TestParser(t *testing.T) {
	tt := []struct {
		name string
		expected
		graph func(t *testing.T) string
	}{
		{name: "can read file", expected: expected{err: false, nodes: 4, edges: 3}, graph: simpleGraph},
		{name: "cant read file, format error", expected: expected{err: true}, graph: invalidGraph},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tmpFile, err := createTmpFile(t, tc.graph(t))
			defer os.Remove(tmpFile.Name())
			if err != nil {
				t.FailNow()
			}
			graph, err := ReadFile(tmpFile.Name())
			if err != nil {
				if tc.expected.err {
					return // short-circuit test,
				}
				fmt.Println("is not format error?")
				t.FailNow()
			}
			nodes := graph.Nodes()
			assert.Equalf(t, nodes.Len(), tc.expected.nodes, "length of nodes: %d does not match expected: %d", nodes.Len(), tc.expected.nodes)
			edges := graph.Edges()
			assert.Equalf(t, edges.Len(), tc.expected.edges, "length of edges: %d does not match expected: %d", edges.Len(), tc.expected.edges)
		})
	}
}

func createTmpFile(t *testing.T, dot string) (*os.File, error) {
	file, err := os.CreateTemp(os.TempDir(), "skii-resort.*.gv")
	if err != nil {
		return nil, errors.Wrap(err, "temporary file creation failed")
	}
	if err := writeDOTToFile(file, dot); err != nil {
		return nil, err
	}

	return file, nil
}

func writeDOTToFile(f *os.File, dot string) error {
	_, err := f.WriteString(dot)
	if err != nil {
		return errors.Wrap(err, "could not write DOT to file")
	}
	if _, err := f.Seek(0, 0); err != nil {
		return errors.Wrap(err, "could not set file pointer to 0,0")
	}
	return nil
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

func invalidGraph(t *testing.T) string {
	const ug = `
graph {
	a--d [color=black]
	a--b [color=red]
	a--c [color=blue]
}---
`
	return ug
}
