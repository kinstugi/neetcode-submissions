func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	ret[0] = 1
	for i := 1; i < n; i++{
		ret[i] = ret[i-1] * nums[i-1]
	}
	suffix := 1
	for i := n-1; i > -1; i--{
		ret[i] *= suffix;
		suffix *= nums[i];
	}
	return ret
}
