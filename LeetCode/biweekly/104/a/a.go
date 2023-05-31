package main

import "strconv"

// https://github.com/Ocyss
func countSeniors(details []string) (ans int) {
	for _, v := range details {
		z, _ := strconv.Atoi(v[14:])
		if z > 60 {
			ans++
		}
	}
	return
}
