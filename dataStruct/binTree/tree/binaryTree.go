package tree

type BinTree struct{
	root *BinTreeNode
}

func NewBinTree(node *BinTreeNode) *BinTree{
	return &BinTree{root: node};
}

func (this *BinTree) Insert(node *BinTreeNode) bool{
	var pre *BinTreeNode;
	var whichChild int;
	
	cur := this.root
	for cur != nil {
		if node.data < cur.data {
			pre = cur;
			cur = cur.lChild
			whichChild = 1;
			continue
		}
		if node.data > cur.data {
			pre = cur;
			cur = cur.rChild
			whichChild = 2;
			continue
		}
		return false;
	}
	if whichChild == 1 {
		pre.lChild = node;
	}
	if whichChild == 2 {
		pre.rChild = node;
	}
	return true;
}

func (this *BinTree) PreOrderTraversal() []int{
	var output []int
	output = this.preOrderTraversal(this.root, output)

	return output
}

func (this *BinTree) preOrderTraversal(root *BinTreeNode, output []int) []int{
	if root == nil {
		return output
	}

	output = append(output, root.data)
	output = this.preOrderTraversal(root.lChild, output);
	output = this.preOrderTraversal(root.rChild, output);

	return output;
}

func (this *BinTree) InOrderTraversal() []int{
	var output []int
	output = this.inOrderTraversal(this.root, output)

	return output
}

func (this *BinTree) inOrderTraversal(root *BinTreeNode, output []int) []int{
	if root == nil {
		return output
	}

	output = this.inOrderTraversal(root.lChild, output);

	output = append(output, root.data)
	output = this.inOrderTraversal(root.rChild, output);

	return output;
}

func (this *BinTree) PostOrderTraversal() []int{
	var output []int
	output = this.postOrderTraversal(this.root, output)

	return output
}

func (this *BinTree) postOrderTraversal(root *BinTreeNode, output []int) []int{
	if root == nil {
		return output
	}

	output = this.postOrderTraversal(root.lChild, output);
	output = this.postOrderTraversal(root.rChild, output);
	output = append(output, root.data)

	return output;
}
