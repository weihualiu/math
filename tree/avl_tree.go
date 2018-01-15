package tree

// avl binary tree
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(tree *Tree) int {
	if tree == nil {
		return -1
	}
	return tree.Height
}

// LL 左旋转
func TreeRetionLL(current *Tree) *Tree {
	tree := new(Tree)
	tree.Copy(current.Left)
	tree.Left = current.Left.Left
	current.Left = tree.Right
	tree.Right = current
	current.Height = max(height(current.Left), height(current.Right)) + 1
	tree.Height = max(height(tree.Left), current.Height) + 1

	return tree
}

// LR
func TreeRetionLR(current *Tree) *Tree {
	current.Left = TreeRetionRR(current.Left)
	return TreeRetionLL(current)
}

// RR  右旋转
func TreeRetionRR(current *Tree) *Tree {
	tree := new(Tree)
	tree.Copy(current.Right)
	tree.Right = current.Right.Right
	current.Right = tree.Left
	tree.Left = current

	current.Height = max(height(current.Left), height(current.Right)) + 1
	tree.Height = max(height(tree.Right), current.Height) + 1

	return tree
}

// RL
func TreeRetionRL(current *Tree) *Tree {
	current.Right = TreeRetionLL(current.Right)
	return TreeRetionRR(current)
}

func AVLTreeAdd(current *Tree, context Context) {
	if current == nil {
		current = TreeRootNew()
		current.Content = context
		current.IsRoot = true
		return
	}
	if current.Content == nil {
		current.Content = context
		return
	}
	tree := new(Tree)
	tree.Content = context
	if current.Content.Compare(context) {
		if current.Left == nil {
			current.Left = tree
		} else {
			AVLTreeAdd(current.Left, context)
		}
	} else {
		if current.Right == nil {
			current.Right = tree
		} else {
			AVLTreeAdd(current.Right, context)
		}
	}
	current = rebalance(current)
}

// 平衡节点
func rebalance(current *Tree) *Tree {
	// 判断树平衡因子BF是否大于1。大于则进行当前节点平衡操作
	leftDeep := TreeDeep(current.Left)
	rightDeep := TreeDeep(current.Right)
	if leftDeep-rightDeep == 2 {
		if current.Left != nil && current.Right == nil && current.Left.Left != nil && current.Left.Right == nil {
			return TreeRetionRR(current)
		} else {
			return TreeRetionRL(current)
		}
	} else if rightDeep-leftDeep == 2 {
		if current.Right != nil && current.Left == nil && current.Right.Right != nil && current.Right.Left == nil {
			return TreeRetionLL(current)
		} else {
			return TreeRetionLR(current)
		}
	}
	return current
}
