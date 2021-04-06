package log

import (
	"strings"
)

func uniqueString(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

func sanitizeLabelKey(key string, isPrefix bool, isValue bool) string {
	if len(key) == 0 {
		return key
	}
	key = strings.TrimSpace(key)
	if len(key) == 0 {
		return key
	}
	if (key[len(key)-1] == ',' || key[len(key)-1] == ':') {
		key = key[:len(key)-1]
	}
	if isPrefix && key[0] >= '0' && key[0] <= '9' {
		key = "_" + key
	}
	if !isValue {
	return strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, key)
    }
    return key
}
