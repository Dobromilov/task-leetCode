func largestAltitude(gain []int) int {
    mx := 0
    temp := 0 
    for i:=0; i<len(gain); i++{
        temp = temp + gain[i]
        mx = max(temp,mx)
    }
    return mx
}