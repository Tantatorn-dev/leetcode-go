package tree

import "leetcode-go/pkg/common"

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*common.TreeNode{root}

	var rets [][]int

	for len(queue) > 0 {
		var ret []int
		currentLevel := len(queue)

		for i := 0; i < currentLevel; i++ {
			popped := queue[0]
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = []*common.TreeNode{}
			}

			ret = append(ret, popped.Val)

			if popped.Left != nil {
				queue = append(queue, popped.Left)
			}

			if popped.Right != nil {
				queue = append(queue, popped.Right)
			}
		}

		rets = append(rets, ret)
	}

	return rets
}
