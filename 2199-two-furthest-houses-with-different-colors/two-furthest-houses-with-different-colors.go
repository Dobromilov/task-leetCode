func maxDistance(colors []int) int {
	n := len(colors)

	i := 0
	for colors[i] == colors[n-1] {
		i++
	}

	j := n - 1
	for colors[j] == colors[0] {
		j--
	}

	if n-1-i > j {
		return n - 1 - i
	}
	return j
}