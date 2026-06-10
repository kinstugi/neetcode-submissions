func maxArea(heights []int) int {
	n := len(heights)
	ans := 0
	for i := 1; i < n; i++{
		for j := 0; j < i; j++{
			area := (i - j) * min(heights[i], heights[j])
			ans = max(area, ans)
		}
	}
	return ans
}
