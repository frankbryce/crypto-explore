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
	_, ans = RsaRun(r, 42, 17, 2753, 61*53)
	if ans != 42 {
		t.Error("42 was not encrypted and decrypted correctly with keys", ans)
	}
}

func Test_Raiser1(t *testing.T) {
	RaiserTest(Raiser1{}, t)
}

func Test_Raiser2(t *testing.T) {
	RaiserTest(Raiser2{}, t)
}

func RaiserBenchmark(r Raiser, k1 uint, k2 uint, mod uint, b *testing.B) {
	for i := 0; i < b.N; i++ {
		RsaRun(r, 42424242, k1, k2, mod)
	}
}

func Benchmark_Raiser1(b *testing.B) {
	RaiserBenchmark(Raiser1{}, 257, 1283585, 1488391, b)
}

func Benchmark_Raiser2(b *testing.B) {
	RaiserBenchmark(Raiser2{}, 257, 1283585, 1488391, b)
}
