func singleNumber(nums []int) int {
    m := make(map[int]int, len(nums)/2+1)
    for i := 0; i<len(nums) ; i++{
        m[nums[i]]++
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