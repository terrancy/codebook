package challenge

//
// UniqueMorseRepresentations
// @Description: 804. 唯一摩尔斯密码词
// @param words
// @return int
//
func UniqueMorseRepresentations(words []string) int {
    list := []string{
        ".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
        "....", "..", ".---", "-.-", ".-..", "--", "-.",
        "---", ".--.", "--.-", ".-.", "...", "-", "..-",
        "...-", ".--", "-..-", "-.--", "--..",
    }
    dict := make(map[string]int)
    res := 0
    for _, word := range words {
        str := ""
        for _, char := range word {
            str += list[char-'a']
        }
        if _, ok := dict[str]; !ok {
            res++
            dict[str] = 1
        }
    }
    return res
}
