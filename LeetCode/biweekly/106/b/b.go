package main

// https://github.com/Ocyss
func longestSemiRepetitiveSubstring(s string) (ans int) {
	l, r := 0, 0
	c := []byte{}
	flag := -1
	f := -1
	for r = range s {
		if len(c) > 0 && c[len(c)-1] == s[r] {
			if flag >= 0 {
				ans = max(ans, r-l)
				l = f + 1
				c = c[flag+1:]
				flag = -1
			} else {
				flag = len(c)
				f = r
			}
		}
		c = append(c, s[r])

	}
	ans = max(ans, r-l+1)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
