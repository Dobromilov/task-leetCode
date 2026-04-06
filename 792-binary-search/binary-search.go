func search(nums []int, target int) int {
    idx, found := slices.BinarySearch(nums, target)
    if found {
        return idx
    }
    return -1
}