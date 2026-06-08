func hasDuplicate(nums []int) bool {
    mp := make(map[int]bool)
    for _, num := range nums{
        if _,ok := mp[num]; ok{
            return true
        }
        mp[num] = true
    }
    return false
}
