//https://leetcode.com/problems/subsets/
//#array #backtracking #bitmanipulation
func subsets(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for st := 0; st < (1 << n); st++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if st&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		ans = append(ans, subset)
	}
	return ans
}