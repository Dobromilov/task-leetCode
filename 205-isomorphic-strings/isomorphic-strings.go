func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m1:= make(map[byte]byte)
    m2:= make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        sChar, tChar := s[i], t[i]

        if val, ok := m1[sChar]; ok && val != tChar{
            return false
        }

        if val, ok := m2[tChar]; ok && val != sChar{
            return false
        }

        m1[sChar] = tChar
        m2[tChar] = sChar
    }
    return true
}