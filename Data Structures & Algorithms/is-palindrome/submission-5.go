func isalnum(ch byte) bool{
	return ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9')
}

func isPalindrome(s string) bool {
	lw := strings.ToLower(s)
	l, r := 0, len(s)-1
	for ; l < r; {
		for ; l < r && !isalnum(lw[l]); {
			l++
		}
		for ; l < r && !isalnum(lw[r]); {
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
