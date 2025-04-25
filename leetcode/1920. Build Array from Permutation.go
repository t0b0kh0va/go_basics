//https://leetcode.com/problems/find-the-maximum-achievable-number/
//#math
func buildArray(nums []int) []int {
	var ans []int
	for _, v := range nums {
		ans = append(ans, nums[v])
	}
	return ans
}