package types

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
)

// Graph - Provides a shim for interaction between the DOT
// unmarshaler and a simple.UndirectedGraph. it also uses
// an id index to quickly locate vertices
type Graph struct {
	*simple.WeightedUndirectedGraph
	idIdx map[string]int64
}

// NewGraph - Return a new Graph with 0 weights and initialized index
func NewGraph() *Graph {
	idx := make(map[string]int64)
	return &Graph{WeightedUndirectedGraph: simple.NewWeightedUndirectedGraph(0, 0), idIdx: idx}
}

// NewEdge - returns a DOT-aware edge.
func (g *Graph) NewEdge(from, to graph.Node) graph.Edge {
	g.addIdsToIndex(from)
	g.addIdsToIndex(to)
	e := g.WeightedUndirectedGraph.NewWeightedEdge(from, to, math.NaN()).(simple.WeightedEdge)
	return &Edge{WeightedEdge: e}
}

// addIdsToIndex - add vertex ID to the index
func (g *Graph) addIdsToIndex(graphNode graph.Node) {
	n := graphNode.(*Node).dotID
	g.idIdx[string(n[:])] = graphNode.ID()
}

// GetAllIndexes - get all indexes
func (g *Graph) GetAllIndexes() map[string]int64 {
	return g.idIdx
}

// NewNode - returns a DOT-aware node.
func (g *Graph) NewNode() graph.Node {
	return &Node{Node: g.WeightedUndirectedGraph.NewNode()}
}

// NewNode - returns a DOT-aware node with DotID.
func (g *Graph) NewNodeWithDotID(dotID []byte) graph.Node {
	return &Node{Node: g.WeightedUndirectedGraph.NewNode(), dotID: dotID}
}

// SetEdge is a shim to allow the DOT unmarshaler to
// add edges to a graph.
func (g *Graph) SetEdge(e graph.Edge) {
	g.WeightedUndirectedGraph.SetWeightedEdge(e.(*Edge))
}

// Edge -is a DOT-aware edge.
type Edge struct {
	simple.WeightedEdge
	Color string
}

// SetAttribute sets the color of the receiver.
func (e *Edge) SetAttribute(attr encoding.Attribute) error {
	if attr.Key != "color" {
		return fmt.Errorf("unable to unmarshal node DOT attribute with key %q", attr.Key)
	}
	e.Color = attr.Value
	e.W = float64(DifficultyConverter(attr.Value))
	return nil
}

// GetColor - Retrieve the color attribute of an edge
func (e *Edge) GetColor() string {
	return e.Color
}

// Node is a DOT-aware node.
type Node struct {
	graph.Node
	dotID []byte
}

// SetDOTID - sets the DOT ID of the node.
func (n *Node) SetDOTID(id string) { n.dotID = []byte(id) }

// String - gets the string from the dotID
func (n *Node) String() string { return string(n.dotID[:]) }

// Bytes - gets the dotID bytes from the node
func (n *Node) Bytes() []byte { return n.dotID }
