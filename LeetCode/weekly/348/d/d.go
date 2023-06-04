package main

// https://github.com/Ocyss
const mod int = 1e9 + 7

func count(x string, y string, minSum int, maxSum int) (ans int) {
	m := len(y)
	n := len(x)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, int, bool) int
	f = func(i, mask, sum int, isNum bool) (res int) {
		if i == m && sum > maxSum {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, sum, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, sum && d == up, true)
			}
		}
		return
	}
	ans = (ans%mod + mod) % mod
	return
}
