func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	ans := 0
	for l < r {
		width := r - l
		area := width * min(heights[l], heights[r])
		ans = max(ans, area)
		if heights[l] < heights[r]{
			l++
		}else{
			r--
		}
	}
	return ans
}
