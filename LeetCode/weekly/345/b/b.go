package main

// https://github.com/Ocyss
func doesValidArrayExist(derived []int) bool {
	n := len(derived)
	if n == 1 {
		if derived[0] == 0 {
			return true
		}
		return false
	} else if n == 2 {
		if derived[0] != derived[1] {
			return false
		}
		return true
	}
	a, b := make([]int, n), make([]int, n)
	if derived[n-1] == 0 {
		a[0] = 0
		a[n-1] = 0
		b[0] = 1
		b[n-1] = 1
	} else {
		a[0] = 0
		a[n-1] = 1
		b[0] = 1
		b[n-1] = 0
	}
	for i := n - 2; i > 0; i-- {
		if derived[i] == 1 {
			a[i] = a[i+1] ^ 1
			b[i] = b[i+1] ^ 1
		} else {
			a[i] = a[i+1]
			b[i] = b[i+1]
		}

	}
	as, bs := 0, 0
	for i, v := range derived {
		if v == a[i] {
			as++
		}
		if v == b[i] {
			bs++
		}
	}
	//fmt.Println(a, as, b, bs)
	if as == n || bs == n {
		return false
	}
	return true
}
