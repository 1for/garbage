package main

import(
	"github.com/1for/garbage/dataStruct/binTree/tree"
	"fmt"
)

func main(){
	header := tree.NewBinTreeNode(2);

	bt := tree.NewBinTree(header);

	bt.Insert(tree.NewBinTreeNode(1));
	bt.Insert(tree.NewBinTreeNode(3));

	fmt.Printf("preorder:%v\n", bt.PreOrderTraversal());
	fmt.Printf("inorder:%v\n", bt.InOrderTraversal());
	fmt.Printf("postOrder:%v\n", bt.PostOrderTraversal());
}
