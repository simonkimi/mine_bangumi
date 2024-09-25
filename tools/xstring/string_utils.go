package xstring

import "strings"

func RemoveEmpty(iter []string) []string {
	result := make([]string, 0, len(iter))
	for _, v := range iter {
		if !IsEmptyOrWhitespace(v) {
			result = append(result, strings.TrimSpace(v))
		}
	}
	return result
}

func IsEmptyOrWhitespace(value string) bool {
	return strings.TrimSpace(value) == ""
}
