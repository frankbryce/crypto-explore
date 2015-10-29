package main

import "testing"

func RaiserTest(r Raiser, t *testing.T) {
	ans := r.Raise(2, 5, 10)
	if ans != 2 {
		t.Error("2^5 (mod 10) should be 2, but it was", ans)
	}
	ans = r.Raise(5, 2, 11)
	if ans != 3 {
		t.Error("5^2 (mod 11) should be 3, but it was", ans)
	}
	ans = r.Raise(4, 5, 12)
	if ans != 4 {
		t.Error("4^5 (mod 12) should be 4, but it was", ans)
	}
	ans = r.Raise(10, 3, 7)
	if ans != 6 {
		t.Error("10^3 (mod 7) should be 6, but it was", ans)
	}
	ans = r.Raise(6, 4, 10)
	if ans != 6 {
		t.Error("6^4 (mod 10) should be 6, but it was", ans)
	}
}

func Test_Raiser1(t *testing.T) {
	RaiserTest(Raiser1{}, t)
}

func Test_Raiser2(t *testing.T) {
	RaiserTest(Raiser2{}, t)
}

func Benchmark_Raiser1(b *testing.B) {
	r := Raiser1{}
	for i := 0; i < b.N; i++ {
		ans := r.Raise(42, 17, 61*53)
		r.Raise(ans, 2753, 61*53)
	}
}

func Benchmark_Raiser2(b *testing.B) {
	r := Raiser2{}
	for i := 0; i < b.N; i++ {
		ans := r.Raise(42, 17, 61*53)
		r.Raise(ans, 2753, 61*53)
	}
}
