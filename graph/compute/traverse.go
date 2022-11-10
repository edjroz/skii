package traverse

import (
	"fmt"
	"skii/graph/types"

	"gonum.org/v1/gonum/graph"
)

func getAllPath(g *types.Graph, startPoint, difficulty string) [][]string {
	path := []string{}
	idxs := g.GetAllIndexes()
	maxDifficulty := types.DifficultyConverter(difficulty)

	return traverse(g, g.Node(idxs[startPoint]), maxDifficulty, path)
}

func traverse(g *types.Graph, node graph.Node, difficulty float64, currentPath []string) [][]string {
	n := node.(*types.Node)
	if isInPath(currentPath, n.String()) {
		return [][]string{currentPath}
	}
	fmt.Printf("Adding %s to path \n", n.String())
	currentPath = append(currentPath, n.String())

	var result [][]string
	neighbors := graph.NodesOf(g.From(node.ID()))
	if len(neighbors) == 0 {
		return [][]string{currentPath} // no more neighbors path is done
	}
	for _, neighbor := range neighbors {
		edge := g.WeightedEdge(node.ID(), neighbor.ID())
		// todo change o debug log
		fmt.Printf("Edge: %s--%s Weight: %f, difficulty: %f\n", n.String(), neighbor.(*types.Node).String(), edge.Weight(), difficulty)

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
