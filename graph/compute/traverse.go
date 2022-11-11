package traverse

import (
	"github.com/edjroz/skii/graph/types"

	"gonum.org/v1/gonum/graph"
)

// GetAllPath - get All paths that can be taken assuming a specific difficulty
// The difficulty is an enum (black, red, blue) the graph has weights assigned to each of the colors black being highest
func GetAllPath(g *types.Graph, startPoint, difficulty string) [][]string {
	path := []string{}
	idxs := g.GetAllIndexes()
	maxDifficulty := types.DifficultyConverter(difficulty)
	// add condition to return empty [][]string on empty graph
	if g.Nodes().Len() <= 1 {
		return [][]string{}
	}
	return traverse(g, g.Node(idxs[startPoint]), maxDifficulty, path)
}

// traverse - Graph traversal, Based on Depth
func traverse(g *types.Graph, node graph.Node, difficulty float64, currentPath []string) [][]string {
	n := node.(*types.Node)
	if isInPath(currentPath, n.String()) {
		return [][]string{currentPath} // We've reached a previous node in the path, gravity doesn't allow cycle without skii lift
	}
	currentPath = append(currentPath, n.String())

	var result [][]string
	neighbors := graph.NodesOf(g.From(node.ID()))
	if len(neighbors) == 0 {
		return [][]string{currentPath} // no more neighbors path is done
	}
	for _, neighbor := range neighbors {
		edge := g.WeightedEdge(node.ID(), neighbor.ID())

		if !isTraversable(edge, difficulty) {
			continue
		}

		result = append(result, traverse(g, neighbor, difficulty, currentPath)...)
	}
	return result
}

func isInPath(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}
func isTraversable(edge graph.WeightedEdge, maxDifficulty float64) bool {
	difficulty := edge.Weight()
	return difficulty <= maxDifficulty
}
