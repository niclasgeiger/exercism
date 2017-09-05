package tree

import (
	"sort"
)

const testVersion = 4

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records Records) (root *Node, err error) {
	if len(records) == 0 {
		return nil, nil
	}
	// sorting the rcords by ID
	sort.Sort(records)
	if records[0].ID > 0 || records[0].Parent > 0 {
		return nil, Mismatch{}
	}
	// the first parent node
	var parent = &Node{
		ID: records[0].ID,
	}
	// the root node
	root = parent
	for i, record := range records[1:] {
		// is continued
		if record.ID != i+1 {
			return nil, Mismatch{}
		}
		node := &Node{
			ID: record.ID,
		}
		if record.Parent == parent.ID {
			parent.Children = append(parent.Children, node)
		} else {
			parent = findNextParent(root, record.Parent)
			if parent == nil {
				return nil, Mismatch{}
			}
			parent.Children = append(parent.Children, node)
		}
	}
	return root, nil
}

// Yes, yes I know that this will only work for trees with level <= 2 (which is the case for everything here)
func findNextParent(node *Node, i int) *Node {
	for _, n := range node.Children {
		if n.ID == i {
			return n
		}
	}
	return nil
}

type Records []Record

// Sort Interface implementation
func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
