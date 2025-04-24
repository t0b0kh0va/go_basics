//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	data := map[int]int{}
	for i, v := range nums {
		k := target - v
		if val, ok := data[k]; ok {
			return []int{val, i}
		}
		data[v] = i
	}
	return nil
}