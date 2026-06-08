import "slices"

func getHash(str string) string{
    s := []rune(str)
    slices.Sort(s)
    return string(s)
}

func groupAnagrams(strs []string) [][]string {
    mp := make(map[string][]string)
    for _, str := range strs{
        hash := getHash(str)
        mp[hash] = append(mp[hash], str)
    }
    ret := make([][]string, len(mp))
    i := 0
    for _, v := range mp{
        ret[i] = v
        i++
    }
    return ret
}
