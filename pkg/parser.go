package pkg

import "strings"

func GetStringInBetweenTwoString(str string, startS string, endS string) (result string, found bool) {
    s := strings.Index(str, startS)
    if s == -1 {
        return result, false
    }
    newS := str[s+len(startS):]
    e := strings.Index(newS, endS)
    if e == -1 {
        return result, false
    }
    result = newS[:e]
    return result, true
}

// TODO: Add checks for code blocks like ```
func CheckHeadline(line string) (string, bool) {
    s := strings.TrimSpace(line)
    if len(line) < 4 {
        return "", false
    }
    if s[0:4] == "####" {
        return "H4", true
    } else if s[0:3] == "###" {
        return "H3", true
    } else if s[0:2] == "##" {
        return "H2", true
    }
    return "", false
}
