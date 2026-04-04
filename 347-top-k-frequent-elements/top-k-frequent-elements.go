func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, v := range nums{
        m[v]++
    }

    res := make([]int, 0, len(m))
    for num := range m {
        res = append(res, num)
    }

    sort.Slice(res, func(i, j int) bool {
        return m[res[i]] > m[res[j]]
    })

    return res[:k]
}