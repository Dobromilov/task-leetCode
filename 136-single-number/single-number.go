func singleNumber(nums []int) int {
    m := make(map[int]int, len(nums)/2+1)
    for i := 0; i<len(nums) ; i++{
        m[nums[i]]++
    }
    for _, n := range nums{
        if m[n] == 1{
            return n
        }
    }
    return 0
}