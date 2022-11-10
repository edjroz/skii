package parser

import (
	"log"
	"os"
	"skii/graph/types"

	"gonum.org/v1/gonum/graph/encoding/dot"
)

func ReadFile(p string) (*types.Graph, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dst := types.NewGraph()
	if err := dot.Unmarshal(b, dst); err != nil {
		return nil, err
	}
	return dst, nil
}
