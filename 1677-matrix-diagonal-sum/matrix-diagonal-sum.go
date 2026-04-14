func diagonalSum(mat [][]int) int {
    n:=len(mat[0])
    sum:=0
    l,r := 0,n-1
    for i:=0; i<n; i++{
        sum = sum + mat[l][r] + mat[l][l]
        r--
        l++
    }
    if n%2==0 {
        return sum
    }else{
        return sum - mat[n/2][n/2]
    }
}