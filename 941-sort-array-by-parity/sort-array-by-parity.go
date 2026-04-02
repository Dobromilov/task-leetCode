func sortArrayByParity(nums []int) []int {
    var ch []int
    var nch []int
    for i:=0; i<len(nums); i++{
        if nums[i]%2==0{
            ch = append(ch,nums[i])
        }else{
            nch = append(nch,nums[i])
        }
    }
    res := append(ch,nch...)
    return res
}