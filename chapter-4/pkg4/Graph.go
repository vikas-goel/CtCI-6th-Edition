package pkg4

type Graph struct {
	directed bool
	edge map[rune][]rune
}

// Initialize a directed or undirected graph, as requested.
func (this Graph) Init(directed bool) *Graph {
	this.directed = directed
	this.edge = make(map[rune][]rune)
	return &this
}

func (this *Graph) AddVertices(vertices ...rune) {
	if this == nil {
		return
	}

	for _, vertex := range vertices {
		if _, ok := this.edge[vertex]; !ok {
			this.edge[vertex] = make([]rune, 0)
		}
	}
}

func (this *Graph) GetVertices() []rune {
	if this == nil {
		return nil
	}

	vertices := make([]rune, 0, len(this.edge))
	for v := range this.edge {
		vertices = append(vertices, v)
	}

	return vertices
}

func (this *Graph) AddEdge(from, to rune) {
	if this == nil {
		return
	}

	// Add both ends to the list of vertices.
	if _, ok := this.edge[from]; !ok {
		this.edge[from] = make([]rune, 0)
	}

	if _, ok := this.edge[to]; !ok {
		this.edge[to] = make([]rune, 0)
	}

	this.edge[from] = append(this.edge[from], to)

	if this.directed {
		this.edge[to] = append(this.edge[to], from)
	}
}

func (this *Graph) GetEdges(source rune) []rune {
	if this == nil {
		return nil
	}

	if _, ok := this.edge[source]; !ok {
		return nil
	}

	return this.edge[source]
}

func (this *Graph) GetAllEdges() map[rune][]rune {
	if this == nil {
		return nil
	}

	return this.edge
}
