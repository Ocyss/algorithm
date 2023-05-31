package main

// https://github.com/Ocyss
func makeSmallestPalindrome(s string) (ans string) {
	n := len(s)
	l, r := 0, n-1
	ss := []byte(s)
	for l < r {
		if ss[l] != ss[r] {
			if ss[l] > ss[r] {
				ss[l] = ss[r]
			} else {
				ss[r] = ss[l]
			}
		}
		l++
		r--
	}
	return string(ss)
}
