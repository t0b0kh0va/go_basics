//https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	var k int
	tmp := x
	for tmp > 0 {
		k = k*10 + tmp%10
		tmp = tmp / 10
	}
	return x == k
}

