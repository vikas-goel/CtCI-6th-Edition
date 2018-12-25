package main

import (
	"fmt"
	graph "./pkg4"
	stack "../chapter-3/stack"
)

type Project = graph.Graph

func BuildOrder(project *Project) {
	if project == nil {
		return
	}

	vertices := project.GetVertices()
	order := stack.Stack{}.New(len(vertices), stack.Runes{})
	visited := make(map[rune]bool)

	var dfs func(rune)
	dfs = func(vertex rune) {
		// Do nothing if already visited this node.
		if _, ok := visited[vertex]; ok {
			return
		}

		visited[vertex] = true
		edges := project.GetEdges(vertex)

		// Push all children in the stack before pushing self.
		for _, e := range edges {
			dfs(e)
		}

		order.Push(vertex)
	}

	for _, v := range vertices {
		dfs(v)
	}

	sequence := make([]rune, 0, len(vertices))
	for !order.Empty() {
		v, _ := order.Pop()
		sequence = append(sequence, v.(rune))
	}

	fmt.Printf("%c\n", sequence)
}

func main() {
	project := Project{}.Init(false)
	for _, v := range [][]rune{
		{'a','b'}, {'f','b'}, {'b','d'}, {'f','a'}, {'d','c'}} {
		project.AddEdge(v[0], v[1])
	}
	project.AddVertices('e')

	BuildOrder(project)
}
