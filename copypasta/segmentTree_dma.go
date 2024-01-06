package copypasta

type SegDma struct {
	left, right *SegDma
	val, add    int
}

func (s *SegDma) Update(start, end, l, r, val int) {
	if l <= start && end <= r {
		s.val += (end - start + 1) * val
		s.add += val
		return
	}
	m := start + (end-start)>>1
	s.pushDown(m-start+1, end-m)
	if l <= m {
		s.left.Update(start, m, l, r, val)
	}
	if r > m {
		s.right.Update(m+1, end, l, r, val)
	}
	s.val = s.left.val + s.right.val
}

func (s *SegDma) Query(start, end, l, r int) (res int) {
	if l <= start && end <= r {
		return s.val
	}
	m := start + (end-start)>>1
	if l <= m {
		res += s.left.Query(start, m, l, r)
	}
	if r > m {
		res += s.right.Query(m+1, end, l, r)
	}
	return
}

func (s *SegDma) pushDown(l, r int) {
	if s.left == nil {
		s.left = &SegDma{}
	}
	if s.right == nil {
		s.right = &SegDma{}
	}
	if s.add == 0 {
		return
	}
	s.left.val += s.add * l
	s.left.add += s.add
	s.right.val += s.add * r
	s.right.add += s.add
	s.add = 0
}
