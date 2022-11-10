package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

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
	dot := fmt.Sprintf("strict graph {\na -- b [color=black]\na -- c [color=red]\na -- d [color=blue]\n}")

	_, err := f.WriteString(dot)
	if err != nil {
		return errors.Wrap(err, "could not write DOT to file")
	}
	if _, err := f.Seek(0, 0); err != nil {
		return errors.Wrap(err, "could not set file pointer to 0,0")
	}
	return nil
}

func TestParser(t *testing.T) {
	tt := []struct {
		name string
	}{
		{name: "can read file"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tmpFile, err := createTmpFile(t)
			defer os.Remove(tmpFile.Name())
			if err != nil {
				t.FailNow()
			}
			graph, err := ReadFile(tmpFile.Name())
			if err != nil {
				t.FailNow()
			}
			nodes := graph.Nodes()
			if nodes.Len() != 4 {
				t.FailNow()
			}
			edges := graph.Edges()
			if edges.Len() != 3 {
				t.FailNow()
			}
		})
	}
}
