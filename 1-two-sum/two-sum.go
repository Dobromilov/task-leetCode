func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums{
        search := target - v
        if temp, ex := m[search] ; ex {
            return []int{temp, i}
        }
        m[v] = i
    }
    return nil
}