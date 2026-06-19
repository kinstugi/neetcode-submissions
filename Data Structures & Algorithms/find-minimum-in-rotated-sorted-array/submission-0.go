func findMin(nums []int) int {
	mn := nums[0]
	for i := range nums {
		if mn > nums[i]{
			mn = nums[i]
		} 
	}
	return mn
}
