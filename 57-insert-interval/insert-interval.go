func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{newInterval}
    }

    leftIdx := n
    l, r := 0, n-1
    for l <= r {
        m := l + (r-l)/2
        if intervals[m][1] >= newInterval[0] {
            leftIdx = m
            r = m - 1
        } else {
            l = m + 1
        }
    }

    rightIdx := -1
    l, r = 0, n-1 
    for l <= r {
        m := l + (r-l)/2
        if intervals[m][0] <= newInterval[1] {
            rightIdx = m
            l = m + 1
        } else {
            r = m - 1
        }
    }

    var res [][]int

    res = append(res, intervals[:leftIdx]...)

    if leftIdx > rightIdx {
        res = append(res, newInterval)
    } else {
        start := min(newInterval[0], intervals[leftIdx][0])
        end := max(newInterval[1], intervals[rightIdx][1])
        res = append(res, []int{start, end})
    }
    res = append(res, intervals[rightIdx+1:]...)

    return res
}

func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
