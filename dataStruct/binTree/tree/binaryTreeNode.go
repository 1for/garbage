package tree

type BinTreeNode struct {
	data int
	lChild *BinTreeNode
	rChild *BinTreeNode
}

func NewBinTreeNode(data int) *BinTreeNode{
	return &BinTreeNode{data: data}
}

func (this *BinTreeNode) SetLChild(node *BinTreeNode) *BinTreeNode{
	oldNode := this.lChild;

	this.lChild = node;

	return oldNode;
}

func (this *BinTreeNode) SetRChild(node *BinTreeNode) *BinTreeNode{
	oldNode := this.rChild;

	this.rChild = node;

	return oldNode;
}

func (this *BinTreeNode) GetLChild(node *BinTreeNode) *BinTreeNode{
	return node.lChild;
}

func (this *BinTreeNode) GetRChild(node *BinTreeNode) *BinTreeNode{
	return node.rChild;
}
