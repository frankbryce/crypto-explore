package main

type Raiser interface {
	Raise(val uint, exp uint, mod uint) (ans uint)
}

type Raiser1 struct {
}

func (r Raiser1) Raise(val uint, exp uint, mod uint) (ans uint) {
	ans = 1
	for i := uint(0); i < exp; i += 1 {
		ans *= val
		ans %= mod
	}
	return ans
}

type Raiser2 struct {
}

func (r Raiser2) Raise(val uint, exp uint, mod uint) (ans uint) {
	ans = 1
	maxVal := (^uint(0))/val - 1
	var valPows []uint = make([]uint, 0, 8)

	for val < maxVal && val != 0 {
		valPows = append(valPows, val)
		val *= val
	}

	c := len(valPows) - 1
	for i := uint(1) << uint(c); i > 0 && exp > 0; {
		if i <= exp {
			ans *= valPows[c]
			ans %= mod
			exp -= i
		} else {
			c -= 1
			i >>= 1
		}
	}

	return ans
}
