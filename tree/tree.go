package tree

// 二叉树存储结构

type Context interface {
	Compare(Context) bool // true 表示当前值大于参数值  false表示当前值小于参数值
	Equals(Context) bool  // true 表示该对象指定值与参数对象指定值一致 false则表示不一致
}

type Tree struct {
	IsRoot  bool
	Left    *Tree
	Right   *Tree
	Content Context
	Height  int
}

func (this *Tree) Copy(src *Tree) {
	this.Left = src.Left
	this.Right = src.Right
	this.Content = src.Content
}

func TreeRootNew() *Tree {
	return &Tree{IsRoot: true}
}

func TreeAdd(current *Tree, content Context) error {
	if current.IsRoot && current.Content == nil {
		current.Content = content
	} else if current.Content != nil {
		tree := new(Tree)
		tree.Content = content
		if current.Content.Compare(content) {
			if current.Left == nil {
				current.Left = tree
			} else {
				return TreeAdd(current.Left, content)
			}
		} else {
			if current.Right == nil {
				current.Right = tree
			} else {
				return TreeAdd(current.Right, content)
			}
		}
	} else {
		panic("tree add object error!")
	}

	return nil
}

// 计算树的深度
func TreeDeep(current *Tree) int {
	if current == nil {
		return 0
	}
	if TreeDeep(current.Left) > TreeDeep(current.Right) {
		return TreeDeep(current.Left) + 1
	} else {
		return TreeDeep(current.Right) + 1
	}

	return 0
}

// 在树中查找数据
func TreeFind(current *Tree, context Context) Context {
	if current == nil || current.Content == nil {
		return nil
	}
	if current.Content.Equals(context) {
		return current.Content
	}
	left := TreeFind(current.Left, context)
	right := TreeFind(current.Right, context)
	if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}

// 删除树中某个节点
//  1、是叶子节点  2、只有左孩子节点  3、只有右孩子节点 4、有两个孩子节点
func TreeDelete(current *Tree, content Context) int {
	var delcount int
	if current.Left != nil && current.Left.Content.Equals(content) {
		current.Left = deleteValid(current.Left, true)
		delcount++
	} else if current.Left != nil {
		delcount += TreeDelete(current.Left, content)
	}
	if current.Right != nil && current.Right.Content.Equals(content) {
		current.Right = deleteValid(current.Right, false)
		delcount++
	} else if current.Right != nil {
		delcount += TreeDelete(current.Right, content)
	}
	if current.IsRoot && current.Content.Equals(content) {
		if current.Left == nil && current.Right == nil {
			current = nil
		} else {
			current = deleteValid(current, true)
			current.IsRoot = true
		}
		delcount++
	}
	return delcount
}

func deleteValid(current *Tree, direct bool) *Tree {
	// 判断是否是叶子节点
	if current.Left == nil && current.Right == nil {
		return nil
	}
	// 是只有左孩子
	if current.Left != nil && current.Right == nil {
		return current.Left
	}
	// 是只有右孩子
	if current.Right != nil && current.Left == nil {
		return current.Right
	}
	// 有左右两孩
	var replace *Tree
	// 查找中序节点，该节点子孙中最接近该值的节点作为当前主节点
	// 如果当前节点是左节点则从子孙节点中选取最大的节点（从左节点中找最右节点作为替代节点）
	if direct {
		replace = TreeFindMax(current.Left, nil)
		// 删除当前节点下的replace节点
		TreeDelete(current, replace.Content)
		replace.Left = current.Left
		replace.Right = current.Right
	} else {
		// 如果是当前是右节点则从子孙中选取最小的节点（从右节点中找最左节点作为替代节点）
		replace = TreeFindMin(current.Right, nil)
		// 删除当前节点下的replace节点
		TreeDelete(current, replace.Content)
		replace.Left = current.Left
		replace.Right = current.Right
	}
	return replace
}

// 从树中找出最右叶子节点
func TreeLeafRight(current *Tree) *Tree {
	if current.Right == nil {
		return current
	}
	return TreeLeafRight(current.Right)
}

// 从树中找出最左叶子节点
func TreeLeafLeft(current *Tree) *Tree {
	if current.Left == nil {
		return current
	}
	return TreeLeafLeft(current.Left)
}

// 从当前节点下找出最大的节点
func TreeFindMax(current *Tree, max *Tree) *Tree {
	if current == nil {
		return max
	}
	if max == nil {
		left := TreeFindMax(current.Left, current)
		right := TreeFindMax(current.Right, current)
		if left.Content.Compare(right.Content) {
			return left
		} else {
			return right
		}
	}
	var tmp *Tree
	if current.Content.Compare(max.Content) {
		tmp = current
	} else {
		tmp = max
	}
	left := TreeFindMax(current.Left, tmp)
	right := TreeFindMax(current.Right, tmp)
	if left.Content.Compare(right.Content) {
		return left
	} else {
		return right
	}
}

// 从当前节点下找出最小的节点
func TreeFindMin(current *Tree, min *Tree) *Tree {
	if current == nil {
		return min
	}
	if min == nil {
		left := TreeFindMin(current.Left, current)
		right := TreeFindMin(current.Right, current)
		if left.Content.Compare(right.Content) {
			return right
		} else {
			return left
		}
	}
	var tmp *Tree
	if current.Content.Compare(min.Content) {
		tmp = min
	} else {
		tmp = current
	}
	left := TreeFindMin(current.Left, tmp)
	right := TreeFindMin(current.Right, tmp)
	if left.Content.Compare(right.Content) {
		return right
	} else {
		return left
	}
}

// 树的宽度
func TreeWidth(current *Tree) int {
	counts := make(map[int]int)
	width(current, 1, counts)
	var tmp int
	for _, v := range counts {
		if tmp < v {
			tmp = v
		}
	}
	return tmp
}

func width(current *Tree, level int, counts map[int]int) {
	if current == nil || current.Content == nil {
		return
	}
	counts[level]++
	width(current.Left, level+1, counts)
	width(current.Right, level+1, counts)
}

// 树的所有节点数
func TreeNodeNum(current *Tree) int {
	if current == nil || current.Content == nil {
		return 0
	}
	if current.Left == nil && current.Right == nil && current.Content != nil {
		return 1
	}
	return nodeNum(current) + 1
}

func nodeNum(current *Tree) int {
	return TreeNodeNum(current.Left) + TreeNodeNum(current.Right)
}

// 树某层中的节点数
func TreeLayerNodeNum(current *Tree, layer int) int {
	var tmp int
	layerNodeNum(current, layer, 1, &tmp)
	return tmp
}

func layerNodeNum(current *Tree, layer int, level int, count *int) {
	if current == nil || current.Content == nil {
		return
	}
	if level == layer {
		*count++
		return
	}
	layerNodeNum(current.Left, layer, level+1, count)
	layerNodeNum(current.Right, layer, level+1, count)
}

// 树的叶子节点数
func TreeLeafNum(current *Tree) int {
	if current == nil || current.Content == nil {
		return 0
	}
	if current.Left == nil && current.Right == nil {
		return 1
	}
	return TreeLeafNum(current.Left) + TreeLeafNum(current.Right)
}

// 树的最大距离（树的直径）
func TreeMaxDistance(current *Tree) int {
	return 0
}

// 树的重心
func TreeWeight(current *Tree) Context {
	return nil
}

// 树中某个节点到根节点的路径
func TreeNodePath(current *Tree, content Context) []Context {
	return nil
}

// 树中两个节点最近的公共父节点
func TreeNodeParent(current *Tree, one, two Context) Context {
	return nil
}

// 树中两个节点间的路径
func TreeNodeFromPath(current *Tree, one, two Context) []Context {
	return nil
}

// 树中两个节点的距离
func TreeNodeFromDistance(current *Tree, one, two Context) int {
	return 0
}

// 翻转树
func TreeRevert(current *Tree) {

}

// 判断树是否是完全二叉树
func TreeIsBTree(current *Tree) bool {
	return false
}

// 判断树是否是满二叉树
func TreeIsFullBTree(current *Tree) bool {
	return false
}

// 判断树是否是平衡二叉树
func TreeIsAVLBTree(current *Tree) bool {
	return false
}
