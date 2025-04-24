//https://leetcode.com/problems/roman-to-integer
//#math #hash_table #string

func romanToInt(s string) int {
	data := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	res := 0
	l := len(s)
	for i := range l {
		if i+1 < l && data[string(s[i])] < data[string(s[i+1])] {
			res -= data[string(s[i])]
		} else {
			res += data[string(s[i])]
		}
	}
	return res
}