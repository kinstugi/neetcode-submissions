func isAnagram(s string, t string) bool {
    mp := make([]int, 26)
    for _, ch := range s{
        mp[ch - 'a'] += 1
    }
    for _, ch := range t{
        mp[ch - 'a'] -= 1
    }
    for _, num := range mp{
        if num != 0{
            return false
        }
    }
    return true
}
