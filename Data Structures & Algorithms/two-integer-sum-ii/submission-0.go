func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for ; l < r; {
		tot := numbers[l] + numbers[r]
		if tot == target{
			return []int{l+1, r+1}
		}
		if tot < target{
			l++
		}else{
			r--
		}
	}
	return []int{-1, -1}
}
