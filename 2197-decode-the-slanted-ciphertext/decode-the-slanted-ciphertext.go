func decodeCiphertext(encodedText string, rows int) string {
    cols := len(encodedText) / rows
    result := make([]byte, 0)
    for j := 0; j < cols; j++ {
        row, col := 0, j
        for row < rows && col < cols {
            index := row*cols + col
            result = append(result, encodedText[index])
            row++
            col++
        }
    }
    return strings.TrimRight(string(result), " ")
}