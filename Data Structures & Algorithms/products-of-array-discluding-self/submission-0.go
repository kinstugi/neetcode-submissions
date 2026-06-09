func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	post := make([]int, n)

	for i, j := 0, n-1; i < n; i, j = i+1, j-1{
		if i == 0{
			pre[i] = 1
			post[j] = 1
			continue
		}
		pre[i] = pre[i-1] * nums[i-1]
		post[j] = post[j+1] * nums[j+1]
	}
	for i := 0; i < n; i++{
		pre[i] *= post[i]
	}
	return pre
}
