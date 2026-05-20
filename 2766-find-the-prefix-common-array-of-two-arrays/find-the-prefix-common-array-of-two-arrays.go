func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)
	seen := make([]int, n+1) 
	matches := 0

	for i := 0; i < n; i++ {
		seen[A[i]]++
		if seen[A[i]] == 2 {
			matches++
		}
		
		seen[B[i]]++
		if seen[B[i]] == 2 {
			matches++
		}
		
		res[i] = matches
	}

	return res
}