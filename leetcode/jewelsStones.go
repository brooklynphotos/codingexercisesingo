// https://leetcode.com/problems/jewels-and-stones/

package leetcode

func numJewelsInStones(J string, S string) int {
	// convert to a map
	JMap := make(map[rune]bool)
	for _, c := range J {
		JMap[c] = true
	}
	count := 0
	for _, c := range S {
		if JMap[c] {
			count++
		}
	}
	return count
}
