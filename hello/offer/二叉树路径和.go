package offer

var tmp []int
var res [][]int
var target1 int

func pathSum(root *TreeNode, target int) [][]int {
	target1 = target
	res = [][]int{}
	tmp = []int{}
	if root == nil {
		return res
	}
	dfs2(root, 0)
	return res
}

func dfs2(node *TreeNode, sum int) {
	sum += node.Val
	tmp = append(tmp, node.Val)
	if node.Left == nil && node.Right == nil && target1 == sum {
		res = append(res, append([]int{}, tmp...))
		return
	}
	if node.Left != nil {
		dfs2(node.Left, sum)
		tmp = tmp[:len(tmp)-1]
	}
	if node.Right != nil {
		dfs2(node.Right, sum)
		tmp = tmp[:len(tmp)-1]
	}

}
