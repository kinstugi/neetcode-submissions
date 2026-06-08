func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    for i, num := range nums{
        dif := target - num
        if idx, ok := mp[dif]; ok{
            return []int {idx, i}
        }
        mp[num] = i
    }
    return []int {-1, -1}
}
