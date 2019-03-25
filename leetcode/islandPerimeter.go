// https://leetcode.com/problems/island-perimeter/

package leetcode

func islandPerimeter(grid [][]int) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			// edges
			if i == 0 || grid[i-1][j] == 0 {
				count++
			}
			if j == 0 || grid[i][j-1] == 0 {
				count++
			}
			if i == len(grid)-1 || grid[i+1][j] == 0 {
				count++
			}
			if j == len(grid[i])-1 || grid[i][j+1] == 0 {
				count++
			}
		}
	}
	return count
}
