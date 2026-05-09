func singleNumber(nums []int) int {
    m := make(map[int]int, len(nums)/2+1)
    for _, n := range nums {
        m[n] += 1
    }
    res:=0
    for _, n := range nums{
        v, _ := m[n]
        if v == 1 {
            res = n
            break
        }
    }
    return res
}