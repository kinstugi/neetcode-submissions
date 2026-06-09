import "slices"

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	fmt.Println(nums)
	ret, cnt, n := 1, 1, len(nums)
	if n < 1 {
		return 0
	}
	for i := 1; i < n; i++{
		if nums[i]-1 == nums[i-1]{
			cnt++
		}else if nums[i] == nums[i-1]{

		}else{
			ret = max(ret, cnt)
			cnt = 1
		}
	}
	return max(ret, cnt)
}
