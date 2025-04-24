//https://leetcode.com/problems/longest-common-prefix/
//#string #trie
func longestCommonPrefix(strs []string) string {
	ans := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(ans) && ans[i] == s[i]; i++ {
		}
		ans = ans[:i]
	}
	return ans
}