//https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
//#array

func canBeIncreasing(nums []int) bool {
	change := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			if change {
				return false
			}
			change = true
			if i == len(nums)-1 {
				return true
			}
			if nums[i+1] > nums[i-1] {
				i++
			} else if i != 1 && nums[i] <= nums[i-2] {
				return false
			}
		}
	}
	return true
}