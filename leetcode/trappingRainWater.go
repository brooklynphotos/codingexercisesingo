package leetcode

/**
 * https://leetcode.com/problems/trapping-rain-water/
 */
func trap(height []int) int {
	total := 0
	if height == nil || len(height) < 3 {
		return total
	}
	leftMax, rightMax, leftPointer, rightPointer := 0, 0, 0, len(height)-1
	for leftPointer < rightPointer {
		// we go left to right as long as left is smaller than right, making the right a bound
		lValue := height[leftPointer]
		rValue := height[rightPointer]
		if lValue < rValue {
			// if we have a new max, we update the leftMax value
			if lValue > leftMax {
				leftMax = lValue
			} else {
				total += leftMax - lValue
			}
			leftPointer++
		} else {
			// same checks
			if rValue > rightMax {
				rightMax = rValue
			} else {
				total += rightMax - rValue
			}
			rightPointer--
		}
	}
	return total
}
