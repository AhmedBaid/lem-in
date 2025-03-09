package graph

import (
	"container/list"
	"strings"

	"lem-in/utils"
)

func BreadthFirstSearch(colony *utils.AntFarm) [][]string {
	const notFound = "not_found"
	var paths [][]string
	startRoom := colony.Start.Name
	endRoom := colony.End.Name

	tree := buildGraph(colony)
	rootNode, rootExists := tree[startRoom]
	if !rootExists {
		return paths
	}

	q := list.New()
	q.PushBack(rootNode)

	parents := make(map[string]string)
	parents[startRoom] = ""

	for q.Len() > 0 {
		currentNode := q.Front().Value.(utils.Node)
		q.Remove(q.Front())

		if strings.EqualFold(currentNode.Value, endRoom) {
			var path []string
			for len(currentNode.Value) > 0 {
				path = append([]string{currentNode.Value}, path...)
				currentNode.Value = parents[currentNode.Value]
			}
			paths = append(paths, path)
		}

		for _, neighbor := range currentNode.Neighbors {
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentNode.Value
				q.PushBack(tree[neighbor])
			}
		}
	}

	return paths
}

func buildGraph(colony *utils.AntFarm) map[string]utils.Node {
	tree := make(map[string]utils.Node)
	for name, _ := range colony.Rooms {
		node := utils.Node{Value: name}
		for _, neighbor := range colony.Links[name] {
			node.Neighbors = append(node.Neighbors, neighbor)
		}
		tree[name] = node
	}
	return tree
}
