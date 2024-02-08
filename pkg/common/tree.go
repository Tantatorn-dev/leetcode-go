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

func maxDepth(root *TreeNode) int {
	if root != nil {
		left := 1 + maxDepth(root.Left)
		right := 1 + maxDepth(root.Right)

		if left > right {
			return left
		} else {
			return right
		}
	}

	return 0
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := 1 + minDepth(root.Left)

	right := 1 + minDepth(root.Right)

	if left == 1 {
		return right
	} else if right == 1 {
		return left
	}

	return min(left, right)
}

func preorderTraversal(root *TreeNode) []int {
	var ret []int

	if root == nil {
		return []int{}
	}

	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)

	ret = append([]int{root.Val}, ret...)

	return ret
}
