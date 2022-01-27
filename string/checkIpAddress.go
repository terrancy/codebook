package strings

import (
    "strconv"
    "strings"
)

func CheckIpAddress(IP string) string {
    // 检查是否是IPv4
    IPv4 := strings.Split(IP, ".")
    flag := true
    if len(IPv4) == 4 {
        for _, val := range IPv4 {
            if !checkIp4Validate(val) {
                flag = false
                break
            }
        }
        if flag == true {
            return "IPv4"
        }
    }
    // 检查是否是IPv6
    IPv6 := strings.Split(IP, ":")
    flag = true
    if len(IPv6) == 8 {
        for _, val := range IPv6 {
            if !checkIp6Validate(val) {
                flag = false
                break
            }
        }
        if flag == true {
            return "IPv6"
        }
    }
    return "Neither"
}

//
//  checkIp4Validate
//  @Description: 检查是否是IPv4
//  @param s
//  @return bool
//
func checkIp4Validate(s string) bool {
    n, err := strconv.Atoi(s)
    if err != nil {
        return false
    }
    if n > 255 || n < 0 {
        return false
    }
    if n > 0 && s[0] == '0' {
        return false
    }
    if n == 0 && len(s) > 1 {
        return false
    }
    return true
}

//
//  checkIp6Validate
//  @Description: 检查是否是IPv6
//  @param s
//  @return bool
//
func checkIp6Validate(s string) bool {
    if len(s) > 4 || len(s) == 0 {
        return false
    }
    for i := 0; i < len(s); i++ {
        if (s[i] >= 'g' && s[i] <= 'z') || (s[i] >= 'G' && s[i] <= 'Z') {
            return false
        }
    }
    return true
}