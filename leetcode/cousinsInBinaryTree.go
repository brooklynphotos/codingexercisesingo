/**
 * https://leetcode.com/contest/weekly-contest-124/problems/cousins-in-binary-tree/
 */
package leetcode

type TreeLevel []*TreeNode

var parents = make(map[int]int)

func IsCousins(root *TreeNode, x int, y int) bool {
	parents[root.Val] = 0
	return isCousinsRec(TreeLevel{root}, x, y)
}

func isCousinsRec(tl TreeLevel, x int, y int) bool {
	foundOnes := find(tl, x, y)
	switch len(foundOnes) {
	case 2:
		return parents[y] != parents[x]
	case 1:
		return false
	case 0:
		return isCousinsRec(nextLevel(tl), x, y)
	default:
		return false
	}
}

func nextLevel(tl TreeLevel) (ret TreeLevel) {
	ret = []*TreeNode{}
	for _, tn := range tl {
		if tn.Left != nil {
			ret = append(ret, tn.Left)
			parents[tn.Left.Val] = tn.Val
		}
		if tn.Right != nil {
			ret = append(ret, tn.Right)
			parents[tn.Right.Val] = tn.Val
		}
	}
	return
}

func find(tl TreeLevel, x int, y int) (ret map[int]bool) {
	ret = make(map[int]bool)
	for _, tn := range tl {
		if tn.Val == x || tn.Val == y {
			ret[tn.Val] = true
		}
	}
	return
}
