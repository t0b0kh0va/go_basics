//https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/
//#array #hash_table #greedy
func maxSum(nums []int) int {
	def := nums[0]
	sum := 0
	for i, k := range nums {
		if !slices.Contains(nums[:i], k) && k > 0 {
			sum += k
			def = math.MaxInt
		} else {
			def = max(def, k)
		}
	}
	if def != math.MaxInt {
		return def
	}
	return sum
}