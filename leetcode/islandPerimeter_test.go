package leetcode

import "testing"

func islandPerimeterTest(t *testing.T) {
	expected := 16
	res := islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	})
	if expected != res {
		t.Errorf("%d != %d", expected, res)
	}
}
