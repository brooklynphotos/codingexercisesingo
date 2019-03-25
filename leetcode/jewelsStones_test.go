package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	res := numJewelsInStones("aA", "aAAbbbb")
	if res != 3 {
		t.Errorf("Expected %d but got %d", 3, res)
	}
}
