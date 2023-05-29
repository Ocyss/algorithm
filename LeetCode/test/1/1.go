package main

import "fmt"

func sort(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	p := a[i]
	for i < j {
		for i < j && a[j] >= p {
			j--
		}
		a[i] = a[j]
		for i < j && a[i] <= p {
			i++
		}
		a[j] = a[i]
	}
	a[i] = p
	sort(a, l, i-1)
	sort(a, i+1, r)
}

func main() {
	b := []int{1, 6, 498, 354, 12341, 564561, 321, 45, 12, 0, 12, -9, 654}
	fmt.Println(b)
	sort(b, 0, len(b)-1)
	fmt.Println(b)
}
