func findMin(nums []int) int {
    mn:=nums[0]
    temp:=nums[0]
    for _,val := range nums{
        if val < mn && val<temp{
            mn = val
            break
        }
        temp = val
    }
    
    return mn
}