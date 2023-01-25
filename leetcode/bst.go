package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode //Left.Val <= Val
	Right *TreeNode //Right.Val >= Val
}

func sortedArrayToBST(nums []int) *TreeNode {
	var t TreeNode
	return &t
}
