package main

import "testing"

func Test_raiseToPower1(t *testing.T) {
	ans := raiseToPower(2, 5, 10)
	if ans != 2 {
		t.Error("2^5 (mod 10) should be 2, but it was", ans)
	}
}

func Test_raiseToPower2(t *testing.T) {
	ans := raiseToPower(5, 2, 11)
	if ans != 3 {
		t.Error("5^2 (mod 11) should be 3, but it was", ans)
	}
}

func Test_raiseToPower3(t *testing.T) {
	ans := raiseToPower(4, 5, 12)
	if ans != 4 {
		t.Error("4^5 (mod 12) should be 4, but it was", ans)
	}
}

func Test_raiseToPower4(t *testing.T) {
	ans := raiseToPower(10, 3, 7)
	if ans != 6 {
		t.Error("10^3 (mod 7) should be 6, but it was", ans)
	}
}

func Test_raiseToPower5(t *testing.T) {
	ans := raiseToPower(6, 4, 10)
	if ans != 6 {
		t.Error("6^4 (mod 10) should be 6, but it was", ans)
	}
}
