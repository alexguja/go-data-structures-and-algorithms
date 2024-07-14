package main

type Node struct {
	name string
}

type Edge struct {
	node   *Node
	weight int
}

type Graph struct {
	Nodes []*Node
	Edges map[string][]*Edge // the key is the node name
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight})
}

func (g *Graph) RemoveEdge(n1, n2 string) {
	removeEdge(g, n1, n2)
	removeEdge(g, n2, n1)
}

func removeEdge(g *Graph, m, n string) {
	edges := g.Edges[m]
	r := -1
	for i, edge := range edges {
		if edge.node.name == n {
			r = i
		}
	}

	if r > -1 {
		edges[r] = edges[len(edges)-1]
		g.Edges[m] = edges[:len(edges)-1]
	}
}

func (g *Graph) RemoveNode(name string) {
	r := -1
	for i, n := range g.Nodes {
		if n.name == name {
			r = i
		}
	}
	if r > -1 {
		g.Nodes[r] = g.Nodes[len(g.Nodes)-1] // remove the node
		g.Nodes = g.Nodes[:len(g.Nodes)-1]
	}

	delete(g.Edges, name) // remove the edge from one side
	// remove the edge from the other side
	for n := range g.Edges {
		removeEdge(g, n, name)
	}
}
