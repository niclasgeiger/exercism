package pov

const testVersion = 2

type Graph struct {
	Root  *Node
	Nodes []*Node
}

func New() *Graph {
	return &Graph{
		Nodes: []*Node{},
	}
}

func (g *Graph) AddNode(nodeLabel string) *Node {
	node := &Node{
		Label:    nodeLabel,
		Parents:  []*Node{},
		Children: []*Node{},
	}
	if g.Root == nil {
		g.Root = node
	}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddArc(from, to string) {
	toNode := g.findNode(to)
	fromNode := g.findNode(from)
	if fromNode == nil {
		fromNode = g.AddNode(from)
		if toNode == g.Root {
			g.Root = fromNode
		}
	}
	fromNode.Children = append(fromNode.Children, toNode)
	toNode.Parents = append(toNode.Parents, fromNode)
}

func (g *Graph) ArcList() (out []string) {
	out = []string{}
	out = getArcs(g.Root)
	return out
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {

	g.ArcList()
	new := g.findNode(newRoot)
	switchUntilOldRoot(new, g.Root, Nodes{})
	g.Root = new
	return g
}

func switchUntilOldRoot(node *Node, root *Node, visited Nodes) {
	visited = append(visited, node)
	if node == root {
		return
	}
	for _, parent := range node.Parents {
		if !visited.containsNode(parent) {
			switchDirection(parent, node)
			switchUntilOldRoot(parent, root, visited)
		}
	}
}

func (g *Graph) findNode(label string) *Node {
	for _, node := range g.Nodes {
		if node.Label == label {
			return node
		}
	}
	return nil
}
