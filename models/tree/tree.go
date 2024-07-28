package tree

type Node struct {
	Key    uint64
	Offset int64
}

type Tree struct {
	Root  Node
	Left  *Tree
	Right *Tree
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func (root *Tree) Height() int {
	if root == nil {
		return 0
	}

	return max(root.Left.Height(), root.Right.Height()) + 1
}

func (root *Tree) GetBalanceFactor() int {
	return root.Left.Height() - root.Right.Height()
}

func (root *Tree) Walk(cb func(*Tree)) {
	if root == nil {
		return
	}

	root.Left.Walk(cb)
	cb(root)
	root.Right.Walk(cb)
}

func LeftRotate(a *Tree) *Tree {
	if a == nil {
		return nil
	}

	b := a.Right
	T2 := b.Left

	a.Right = T2
	b.Left = a

	return b
}

func RighRotate(a *Tree) *Tree {
	if a == nil {
		return nil
	}

	b := a.Left
	T2 := b.Right

	b.Right = a
	a.Left = T2

	return b
}

func Insert(root *Tree, node Node) *Tree {
	if root == nil {
		return &Tree{
			Root: node,
		}
	}

	if node.Key > root.Root.Key {
		root.Right = Insert(root.Right, node)
	}

	if node.Key < root.Root.Key {
		root.Left = Insert(root.Left, node)
	}

	// start balancing the tree
	balance_factor := root.GetBalanceFactor()

	// right rotation
	if balance_factor > 1 {
		if node.Key < root.Left.Root.Key {
			root = RighRotate(root)
		}

		if node.Key > root.Left.Root.Key {
			root.Left = LeftRotate(root.Left)
			root = RighRotate(root)
		}
	}

	// left rotation
	if balance_factor < -1 {
		if node.Key > root.Right.Root.Key {
			root = LeftRotate(root)
		}

		if node.Key < root.Right.Root.Key {
			root.Right = RighRotate(root.Right)
			root = LeftRotate(root)
		}
	}

	return root
}
