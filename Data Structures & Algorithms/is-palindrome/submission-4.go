func isalnum(ch rune) bool{
	return (rune('a') <= ch && ch <= rune('z')) || (rune('0') <= ch && ch <= rune('9'))
}

func isPalindrome(s string) bool {
	lw := strings.ToLower(s)
	l, r := 0, len(s)-1
	for ; l < r; {
		for ; l < r && !isalnum(rune(lw[l])); {
			l++;
		}
		for ; l<r && !isalnum(rune(lw[r])); {
			r--
		}
		if lw[r] != lw[l]{
			return false
		}
		l++
		r--
	}
	return true
}
