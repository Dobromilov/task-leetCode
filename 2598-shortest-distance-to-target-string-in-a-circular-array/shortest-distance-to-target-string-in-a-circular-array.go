func closestTarget(words []string, target string, startIndex int) int {
    n:=len(words)
    
    count_l, count_r := 0, 0
    ind_l := startIndex
    ind_r := startIndex
    res_ind1, res_ind2 := n, n 
    for i:=0; i<len(words); i++{
        if words[ind_l] == target {
            res_ind1 = min(res_ind1, count_l)
        }
        ind_l = (ind_l - 1 + n) % n
        count_l++

        if words[ind_r] == target {
            res_ind2 = min(res_ind2, count_r)
        }
        ind_r = (ind_r + 1) % n
        count_r++
    }
    if res_ind1 == n && res_ind2 == n {
        return -1
    }
    
    return min(res_ind1, res_ind2)
}