package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	root := &TreeNode{}
	addNode(root, 1);

	ret := rightSideView(root)
	fmt.Println(ret)
}

/**
 * n 表示 node 数
 */
func addNode(root *TreeNode, n int) {
	if(root == nil){
		return ;
	}
	
	root.Val = n;

	if n < 8 {
		root.Left = &TreeNode{}
		addNode(root.Left, 2*n)

		root.Right = &TreeNode{}	
		addNode(root.Right, 2*n+1)
	}
	return ;
}

func printTree(root *TreeNode){
	
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func rightSideView(root *TreeNode) []int {
	var arr []int
	n := 0
	arr = DLR(root, n, arr)
	
	return arr
}

/**
 * 中序遍历
 */
func DLR(root *TreeNode, n int,  arr []int) []int {
//	fmt.Println(root)
	if root == nil {
		return arr
	}

	if len(arr) < (n+1) {
		arr = append(arr, root.Val)
	}else{
		arr[n] = root.Val
	}
	
	n++
	arr = DLR(root.Left, n, arr)	
	arr = DLR(root.Right, n, arr)

	return arr
}

