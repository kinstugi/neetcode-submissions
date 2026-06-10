func lengthOfLongestSubstring(s string) int {
	l, ans := 0, 0
	n := len(s)
	mp := make([]int, 128)

	for i := 0; i < n; i++{
		for mp[s[i]] > 0 {
			mp[s[l]]--
			l++
		}
		mp[s[i]]++
		ans = max(ans, i - l + 1)
	}
	return ans
}
