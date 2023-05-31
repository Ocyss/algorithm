package main

// https://github.com/Ocyss
func sumOfMultiples(n int) (ans int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			ans += i
		}
	}
	return
}
