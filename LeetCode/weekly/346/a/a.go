package main

import (
	"fmt"
	"strings"
)

// https://github.com/Ocyss
func minLength(s string) int {
	for {
		ab := strings.Index(s, "AB")
		if ab >= 0 {
			s = s[:ab] + s[ab+2:]
		}
		fmt.Println(s)
		cd := strings.Index(s, "CD")
		if cd >= 0 {
			s = s[:cd] + s[cd+2:]
		}
		if cd == ab && ab == -1 {
			break
		}
	}
	return len(s)
}
