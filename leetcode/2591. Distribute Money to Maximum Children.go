//https://leetcode.com/problems/distribute-money-to-maximum-children/
//#math #greedy

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if money > children*8 {
		return children - 1
	}
	ans := 0
	for money > 0 && money-8 >= children-1 {
		ans++
		money -= 8
		children--
	}
	if money == 4 && children == 1 {
		ans--
	}
	return ans
}