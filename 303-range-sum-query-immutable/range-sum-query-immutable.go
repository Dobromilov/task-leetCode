type NumArray struct {
    sum []int
}


func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums)+1)
    for i:=0; i<len(nums); i++{
        prefix[i+1] = prefix[i]+nums[i]
    }
    return NumArray{sum: prefix}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.sum[right+1] - this.sum[left]
}
