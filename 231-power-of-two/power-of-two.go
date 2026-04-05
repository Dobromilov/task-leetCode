func isPowerOfTwo(n int) bool {
    if n<=0 {
        return false
    }
    temp := 1
    for i := 0; i < 31; i++{
        if temp == n{
            return true
        } else if temp > n{
            return false
        } 
        temp<<=1
    }
    return false
}