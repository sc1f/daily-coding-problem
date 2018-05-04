package algorithms

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NextNode struct {
	Value int
	Left  *Node
	Right *Node
	Next  *Node
}

func PreInConstructor(preorder []int, inorder []int) *Node {
	if len(inorder) > 0 {
		ind := 0
		root_val = preorder[0]
		// find the node in the inorder traversal
		for index, num := range inorder {
			if num == root_val {
				ind = index
			}
		}
		// remove from slice via variadic unpacking
		preorder = append(preorder[:1], preorder[2:]...)
		// init new node with nil left/right
		root = Node{inorder[ind], nil, nil}
		// go left with elements left of where we found the node, go right in the same way
		root.Left = PreInConstructor(preorder, inorder[0:ind])
		root.Right = PreInConstructor(preorder, inorder[ind+1:len(inorder)])
		return *root
	}
}

func PostInConstructor(postorder []int, inorder []int) *Node {
	// TODO finish
	if len(inorder) > 0 {
		ind := 0
		root_val = postorder[len(postorder)-1]
		// find index of the element
		for index, num := range inorder {
			if num == root_val {
				ind = index
			}
		}
		// remove from slice via variadic unpacking
		preorder = append(postorder[1:], postorder[:2]...)
		// init new node with nil left/right
		root = Node{inorder[ind], nil, nil}
		root.Right = PostInConstructor(postorder, inorder[ind+1:len(inorder)])
		root.Left = PostInConstructor(postorder, inorder[0:ind])
		return *root
	}
}

func (*Node) ValidateBST(root *Node) bool {
	return validator(root, -2147483648, 2147483647) // assume int32
}

func validator(root *Node, min int32, max int32) bool {
	if root == nil {
		return true
	}

	if root.Value > max || root.Value < min {
		return False
	}

	return validator(root.left, min, root.Value-1) && validator(root.right, root.Value+1, max)
}

func (*NextNode) PopulateNextRight() {
	//
}
