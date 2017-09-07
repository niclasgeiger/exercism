package pov

import "fmt"

const testVersion = 2

type Graph struct {
	Root  *Node
	Nodes []*Node
}

type Node struct {
	Label string
	To    []*Node
}

func New() *Graph {
	node := &Node{
		Label: "parent",
	}
	return &Graph{
		Root: node,
		Nodes: []*Node{
			node,
		},
	}
}

func (g *Graph) AddNode(nodeLabel string) *Node {
	node := &Node{
		Label: nodeLabel,
		To:    []*Node{},
	}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddArc(from, to string) {
	defer func() {
		if recover() != nil {
			fmt.Printf("\nError recover: \n from:%s\n to:%s\n", from, to)
			fmt.Printf("Graph:%s\n\n", g)
		}
	}()
	fromNode := g.findNode(from)
	if fromNode == nil {
		fromNode = g.AddNode(from)
	}
	toNode := g.findNode(to)
	fromNode.To = append(fromNode.To, toNode)
}

func (g *Graph) ArcList() (out []string) {
	out = []string{}
	out = getArcs(g.Root)
	return out
}

func getArcs(node *Node) (out []string) {
	out = []string{}
	if node != nil {
		for _, neighbor := range node.To {
			out = append(out, fmt.Sprintf("%s -> %s", node.Label, neighbor.Label))
			out = append(out, getArcs(neighbor)...)
		}
	}
	return out
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	return new(Graph)
}

func (g *Graph) findNode(label string) *Node {
	for _, node := range g.Nodes {
		if node.Label == label {
			return node
		}
	}
	return nil
}
