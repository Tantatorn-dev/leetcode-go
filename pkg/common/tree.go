package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2

	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = nums[mid]

	if mid > 0 {
		root.Left = SortedArrayToBST(nums[:mid])
		root.Right = SortedArrayToBST(nums[mid+1:])
	}

	return root
}
