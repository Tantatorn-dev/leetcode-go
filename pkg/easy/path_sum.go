package easy

import "leetcode-go/pkg/common"

func hasPathSum(root *common.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
	} else if root.Left != nil {
		return hasPathSum(root.Left, targetSum)
	} else {
		return hasPathSum(root.Right, targetSum)
	}
}
