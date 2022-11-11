package traverse

import (
	"os"
	"testing"

	"github.com/edjroz/skii/graph/parser"
	"github.com/edjroz/skii/graph/types"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type expected struct {
	noOfRoutes int
	edges      [][]string
}

func TestGetAllNodes(t *testing.T) {
	tt := []struct {
		name string
		expected
		graph      func(t *testing.T) string
		startPoint string
		diff       string
	}{
		{name: "base example, blue", expected: expected{noOfRoutes: 1, edges: [][]string{[]string{"a", "c"}}}, graph: simpleGraph, startPoint: "a", diff: "blue"},
		{name: "base example, red", expected: expected{noOfRoutes: 2, edges: [][]string{[]string{"a", "c"}, []string{"a", "b"}}}, graph: simpleGraph, startPoint: "a", diff: "red"},
		{name: "base example, black", expected: expected{noOfRoutes: 3, edges: [][]string{[]string{"a", "d"}, []string{"a", "b"}, []string{"a", "c"}}}, graph: simpleGraph, startPoint: "a", diff: "black"},

		{name: "single Vertice", expected: expected{noOfRoutes: 0}, graph: singleVertex, startPoint: "a", diff: "blue"},
		{name: "no vertex", expected: expected{noOfRoutes: 0}, graph: noVertex, startPoint: "a", diff: "blue"},

		{name: "complex graph, blue, start A", expected: expected{noOfRoutes: 1, edges: [][]string{[]string{"a", "c"}}}, graph: complexGraph, startPoint: "a", diff: "blue"},
		{name: "complex graph, red, start A", expected: expected{noOfRoutes: 4, edges: [][]string{[]string{"a", "c"}, []string{"a", "b"}, []string{"a", "e"}, []string{"a", "e", "f"}}}, graph: complexGraph, startPoint: "a", diff: "red"},
		{
			name: "complex graph, black, start A",
			expected: expected{
				noOfRoutes: 6,
				edges: [][]string{
					[]string{"a", "c"},
					[]string{"a", "b"},
					[]string{"a", "b", "d"},
					[]string{"a", "e"},
					[]string{"a", "e", "f"},
					[]string{"a", "e", "g"},
				},
			}, graph: complexGraph, startPoint: "a", diff: "black"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			g := createGraph(t, tc.graph(t))
			ans := GetAllPath(g, tc.startPoint, tc.diff)
			assert.Equal(t, len(ans), tc.expected.noOfRoutes, "received routes don't match expected rotues")
			if tc.expected.noOfRoutes > 0 {
				return // short circuit
			}
			assert.ObjectsAreEqualValues(ans, tc.expected.edges)
		})
	}

}

func createGraph(t *testing.T, dot string) *types.Graph {
	tmpFile, err := createTmpFile(t, dot)
	defer os.Remove(tmpFile.Name())
	if err != nil {
		t.FailNow()
	}
	graph, err := parser.ReadFile(tmpFile.Name())
	if err != nil {
		t.FailNow()
	}
	return graph
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

func complexGraph(t *testing.T) string {
	const ug = `
	graph {
		a--b [color=red]
		a--c [color=blue]
		b--d [color=black]
		a--e [color=red]
		e--f [color=blue]
		e--g [color=black]
	}	`
	return ug
}

func singleVertex(t *testing.T) string {
	const ug = `
	graph {
		a 
	}	`
	return ug
}

func noVertex(t *testing.T) string {
	const ug = `
	graph {
	}	`
	return ug
}
