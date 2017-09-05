package binarysearchtree

const testVersion = 1

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{
		data: data,
	}
}

func (bst *SearchTreeData) Insert(data int) {
	if data > bst.data {
		if bst.right != nil {
			bst.right.Insert(data)
		} else {
			bst.right = Bst(data)
		}
	} else {
		if bst.left != nil {
			bst.left.Insert(data)
		} else {
			bst.left = Bst(data)
		}
	}
}

func (bst *SearchTreeData) MapString(f func(int) string) (out []string) {
	if bst.left != nil {
		left := bst.left.MapString(f)
		out = append(left, out...)
	}
	out = append(out, f(bst.data))
	if bst.right != nil {
		out = append(out, bst.right.MapString(f)...)
	}
	return out

}

func (bst *SearchTreeData) MapInt(f func(int) int) (out []int) {
	if bst.left != nil {
		left := bst.left.MapInt(f)
		out = append(left, out...)
	}
	out = append(out, f(bst.data))
	if bst.right != nil {
		out = append(out, bst.right.MapInt(f)...)
	}
	return out
}
