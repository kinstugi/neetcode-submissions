import "slices"
type Pair struct{
    count int
    num int
}

func topKFrequent(nums []int, k int) []int {
    mp := make(map[int]int)
    for _,num := range nums{
        mp[num]++
    }
    ret := make([]int, 0, k)
    tup := make([]Pair, 0, len(mp))
    for k, v := range mp{
        tup = append(tup, Pair{v, k})
    }
    slices.SortFunc(tup, func(a , b Pair)int{
        if b.count != a.count{
            return b.count - a.count
        }
        return b.num - a.num
    })
    for i := 0; i < k; i++{
        ret = append(ret, tup[i].num)
    }
    return ret
}
