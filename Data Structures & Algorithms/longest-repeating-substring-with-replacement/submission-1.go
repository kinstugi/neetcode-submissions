func characterReplacement(s string, k int) int {
	mp := make([]int, 26)
	for i := 0; i < len(s); i++{
		d := s[i] - 'A'
		mp[d]++	
	}
	ans := 0
	for i := 0; i < 26; i++{
		if mp[i] == 0 {
			continue
		}
		ch, rem, l := byte('A' + i), k, 0
		for r := 0; r < len(s); r++{
			for rem == 0 && s[r] != ch{
				if s[l] != ch{
					rem++
				}
				l++
			}
			if s[r] != ch {
				rem--
			}
			ans = max(ans, r - l + 1)
		}
	}
	return ans
}
