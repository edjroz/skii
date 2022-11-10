package traverse

import (
	"fmt"
	"os"
	"skii/graph/parser"
	"skii/graph/types"
	"testing"

	"github.com/pkg/errors"
)

func TestGetAllNodes(t *testing.T) {
	startPoint := "f"
	g := createGraph(t)

	ans := getAllPath(g, startPoint, "red")
	fmt.Printf("final answer = %+v\n", ans)
	t.FailNow()
}

func createGraph(t *testing.T) *types.Graph {
	tmpFile, err := createTmpFile(t)
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

func createTmpFile(t *testing.T) (*os.File, error) {
	file, err := os.CreateTemp(os.TempDir(), "skii-resort.*.gv")
	if err != nil {
		return nil, errors.Wrap(err, "temporary file creation failed")
	}
	if err := writeDOTToFile(file); err != nil {
		return nil, err
	}

	return file, nil
}

func writeDOTToFile(f *os.File) error {
	const dot = `
	graph {
		a--b [color=red]
		a--c [color=blue]
		b--d [color=black]
		a--e [color=red]
		e--f [color=blue]
		e--g [color=black]
	}	`

	_, err := f.WriteString(dot)
	if err != nil {
		return errors.Wrap(err, "could not write DOT to file")
	}
	if _, err := f.Seek(0, 0); err != nil {
		return errors.Wrap(err, "could not set file pointer to 0,0")
	}
	return nil
}
