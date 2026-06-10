func maxProfit(prices []int) int {
	mn, n := prices[0], len(prices)
	ans := 0
	for i := 0; i < n; i++{
		mn = min(mn, prices[i])
		ans = max(ans, prices[i] - mn)
	}
	return ans
}
