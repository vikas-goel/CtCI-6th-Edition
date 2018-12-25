package main

import (
	"fmt"
	graph "./pkg4"
)

type Graph = graph.Graph

func Route(graph *Graph, source, target rune) bool {
	if graph == nil {
		return false
	}

	visited := make(map[rune]bool)

	var dfs func(rune) bool
	dfs = func(vertex rune) bool {
		// Do nothing if already visited this node.
		if _, ok := visited[vertex]; ok {
			return false
		}

		visited[vertex] = true
		edges := graph.GetEdges(vertex)

		// Find if there is a direc edge or a path.
		for _, e := range edges {
			if e == target || dfs(e) {
				return true
			}
		}

		return false
	}

	if graph.GetEdges(source) == nil {
		return false
	}

	return dfs(source)
}

func main() {
	graph := Graph{}.Init(false)
	for _, v := range [][]rune{
		{'a','b'}, {'f','b'}, {'b','d'}, {'f','a'}, {'d','c'}} {
		graph.AddEdge(v[0], v[1])
	}

	source, target := 'a', 'e'
	fmt.Printf("(%c, %c) = %v\n", source, target, Route(graph, source, target))
}
