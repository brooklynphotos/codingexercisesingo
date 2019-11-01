package leetcode

import "testing"

func TestTrapping(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{0, 1, 2}, 0},
		{[]int{2, 1, 0}, 0},
	}
	for _, tt := range tests {
		res := trap(tt.input)
		if res != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, res)
		}
	}
}
