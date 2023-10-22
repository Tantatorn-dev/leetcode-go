package easy

import "leetcode-go/pkg/common"

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else if p != nil || q != nil {
		return false
	}

	return true
}
