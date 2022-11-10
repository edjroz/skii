package types

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
)

// Graph provides a shim for interaction between the DOT
// unmarshaler and a simple.UndirectedGraph.
type Graph struct {
	*simple.WeightedUndirectedGraph
	idIdx map[string]int64
}

func NewGraph() *Graph {
	// create map otherwise panic abound
	idx := make(map[string]int64)
	return &Graph{WeightedUndirectedGraph: simple.NewWeightedUndirectedGraph(0, 0), idIdx: idx}
}

// NewEdge returns a DOT-aware edge.
func (g *Graph) NewEdge(from, to graph.Node) graph.Edge {
	g.addIdsToIndex(from)
	g.addIdsToIndex(to)
	e := g.WeightedUndirectedGraph.NewWeightedEdge(from, to, math.NaN()).(simple.WeightedEdge)
	return &Edge{WeightedEdge: e}
}
func (g *Graph) addIdsToIndex(graphNode graph.Node) {
	n := graphNode.(*Node).dotID
	g.idIdx[n] = graphNode.ID()
}
func (g *Graph) GetAllIndexes() map[string]int64 {
	return g.idIdx
}

// NewNode returns a DOT-aware node.
func (g *Graph) NewNode() graph.Node {
	// todo: inject map to keep reference of nodes, if done we can find the starting point quicker
	return &Node{Node: g.WeightedUndirectedGraph.NewNode()}
}

// SetEdge is a shim to allow the DOT unmarshaler to
// add  edges to a graph.
func (g *Graph) SetEdge(e graph.Edge) {
	g.WeightedUndirectedGraph.SetWeightedEdge(e.(*Edge))
}

// Edge is a DOT-aware  edge.
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

func (e *Edge) GetColor() string {
	return e.Color
}

// node is a DOT-aware node.
type Node struct {
	graph.Node
	dotID string
}

// SetDOTID sets the DOT ID of the node.
func (n *Node) SetDOTID(id string) { n.dotID = id }

func (n *Node) String() string { return n.dotID }
