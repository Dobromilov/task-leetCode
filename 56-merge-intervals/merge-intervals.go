func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    
    res := make([][]int, len(intervals))
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    x := 0
    res[x] = intervals[0]
    
    for i := 1; i < len(intervals); i++ {
        if res[x][1] >= intervals[i][0] {
            if intervals[i][1] > res[x][1] {
                res[x][1] = intervals[i][1]
            }
        } else {
            x++
            res[x] = intervals[i]
        }
    }
    
    return res[:x+1]
}