import "slices"

func threeSum(nums []int) [][]int {
    n := len(nums)
    slices.Sort(nums)
    var ret [][]int

    for i:= 0; i < n-2; i++{
        l, r := i+1, n-1
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        for l < r {
            tot := nums[i] + nums[l] + nums[r]
            if tot == 0{
                ret = append(ret, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1]{
                    l++
                }
                for l < r && nums[r] == nums[r-1]{
                    r--
                }
                l++
                r--
            }else if tot < 0{
                l++
            }else{
                r--
            }
        }
    }
    return ret
}
