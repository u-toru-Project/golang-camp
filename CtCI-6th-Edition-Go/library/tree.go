package library

type TreeNode struct {
	Data   int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	Size   int
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{Data: data, Size: 1}
}

func (n *TreeNode) SetLeftChild(left *TreeNode) {
	n.Left = left
	if left != nil {
		left.Parent = n
	}
}

func (n *TreeNode) SetRightChild(right *TreeNode) {
	n.Right = right
	if right != nil {
		right.Parent = n
	}
}

func (n *TreeNode) InsertInOrder(data int) {
	if data <= n.Data {
		if n.Left == nil {
			n.SetLeftChild(NewTreeNode(data))
		} else {
			n.Left.InsertInOrder(data)
		}
	} else if n.Right == nil {
		n.SetRightChild(NewTreeNode(data))
	} else {
		n.Right.InsertInOrder(data)
	}
	n.Size++
}

func (n *TreeNode) IsBST() bool {
	if n.Left != nil && (n.Data < n.Left.Data || !n.Left.IsBST()) {
		return false
	}
	if n.Right != nil && (n.Data >= n.Right.Data || !n.Right.IsBST()) {
		return false
	}
	return true
}

func (n *TreeNode) Height() int {
	left, right := 0, 0
	if n.Left != nil {
		left = n.Left.Height()
	}
	if n.Right != nil {
		right = n.Right.Height()
	}
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func (n *TreeNode) Find(data int) *TreeNode {
	if data == n.Data {
		return n
	}
	if data <= n.Data {
		if n.Left == nil {
			return nil
		}
		return n.Left.Find(data)
	}
	if n.Right == nil {
		return nil
	}
	return n.Right.Find(data)
}

func CreateMinimalBST(array []int) *TreeNode {
	return createMinimalBSTRange(array, 0, len(array)-1)
}

func createMinimalBSTRange(array []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	node := NewTreeNode(array[mid])
	node.SetLeftChild(createMinimalBSTRange(array, start, mid-1))
	node.SetRightChild(createMinimalBSTRange(array, mid+1, end))
	return node
}
