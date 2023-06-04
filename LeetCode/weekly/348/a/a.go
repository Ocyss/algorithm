package main

// https://github.com/Ocyss
func minimizedStringLength(s string) (ans int) {
	set := map[byte]bool{}
	for i := range s {
		set[s[i]] = true
	}
	return len(set)
}
