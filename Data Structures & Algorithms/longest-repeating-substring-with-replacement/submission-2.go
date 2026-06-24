func characterReplacement(s string, k int) int {
	counts := make([]int, 26)
	maxFreq, left, maxLen := 0, 0, 0
	
	for  right := 0; right < len(s); right++{
		charIdx := s[right] - 'A'
		counts[charIdx]++
		maxFreq = max(maxFreq, counts[charIdx])

		if (right - left + 1) - maxFreq > k {
			counts[s[left]-'A']--
			left++
		}
		maxLen = max(maxLen, right - left + 1)
	}
	return maxLen
}
