func decodeCiphertext(encodedText string, rows int) string {
    if rows == 0 {
        return ""
    }
    cols := len(encodedText) / rows
    matrix := make([][]rune, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = make([]rune, cols)
    }
    count := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            matrix[i][j] = rune(encodedText[count])
            count++
        }
    }

    var res []rune
    for j := 0; j < cols; j++ {
        for i := 0; i < rows && i+j < cols; i++ {
            res = append(res, matrix[i][i+j])
        }
    }
    return strings.TrimRight(string(res), " ")
}