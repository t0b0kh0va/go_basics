//https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/
//#array #slidingwindow

func longestAlternatingSubarray(nums []int, threshold int) int {
	longest := 0
	l := 0
	for i := range nums {
		if nums[i] > threshold {
			l = 0
		} else if l&1 != nums[i]&1 {
			l = l & 1
		} else {
			l += 1
			longest = max(longest, l)
		}
	}
	return longest
}