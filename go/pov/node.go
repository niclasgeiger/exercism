package pov

import "fmt"

type Node struct {
	Label    string
	Parents  []*Node
	Children []*Node
}

func switchDirection(parent, child *Node) {
	parent.DropChildren(child)
	parent.Parents = append(parent.Parents, child)
	child.DropParent(parent)
	child.Children = append(child.Children, parent)
}

func (n *Node) DropChildren(drop *Node) {
	index := -1
	for i, to := range n.Children {
		if to == drop {
			index = i
		}
	}
	if index < 0 {
		return
	}
	if index == 0 {
		n.Children = n.Children[1:]
		return
	}
	if index == len(n.Children)-1 {
		n.Children = n.Children[:len(n.Children)-1]
		return
	}
	n.Children = append(n.Children[:index], n.Children[index+1:]...)
}

func (n *Node) DropParent(drop *Node) {
	index := -1
	for i, to := range n.Parents {
		if to == drop {
			index = i
		}
	}
	if index < 0 {
		return
	}
	if index == 0 {
		n.Parents = n.Parents[1:]
		return
	}
	if index == len(n.Parents)-1 {
		n.Parents = n.Parents[:len(n.Parents)-1]
		return
	}
	n.Parents = append(n.Parents[:index], n.Parents[index+1:]...)
}

func getArcs(node *Node) (out []string) {
	out = []string{}
	if node != nil {
		for _, children := range node.Children {
			out = append(out, fmt.Sprintf("%s -> %s", node.Label, children.Label))
			out = append(out, getArcs(children)...)
		}
	}
	return out
}

type Nodes []*Node

func (list Nodes) containsNode(node *Node) bool {
	for _, ele := range list {
		if node == ele {
			return true
		}
	}
	return false
}
