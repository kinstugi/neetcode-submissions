func isValid(s string) bool {
    stk := make([]byte, 0, len(s))
	for c := range s {
		if s[c] == '(' || s[c] == '[' || s[c] == '{'{
			stk = append(stk, s[c])
		}else if len(stk) < 1{
			return false
		}else{
			last := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if s[c] - last > 2 {
				return false
			}
		}
	}
	return len(stk) == 0
}
