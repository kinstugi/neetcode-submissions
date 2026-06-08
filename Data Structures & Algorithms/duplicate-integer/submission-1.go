func hasDuplicate(nums []int) bool {
    mp := make(map[int]int)
    for _, num := range nums{
        if _,ok := mp[num]; ok{
            return true
        }
        mp[num] = 1
    }
    return false
}
