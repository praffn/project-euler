package util

import "strings"

func IsPalindrome(str string) bool {
	split := strings.Split(str, "")
	for start, end := 0, len(split); start < len(split) && start != end; start++ {
		end--
		if split[start] != split[end] {
			return false
		}
	}
	return true
}
