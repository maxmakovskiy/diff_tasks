package main

import "testing"

func TestDigitPow(t *testing.T) {
	tables := []struct {
		n, p, res int
	}{
		{89, 1, 1},
		{92, 1, -1},
		{695, 2, 2},
		{46288, 3, 51},
		{114, 3, 9},
	}

	for _, table := range tables {
		digPowResult := DigitPow(table.n, table.p)
		if digPowResult != table.res {
			t.Errorf("DigitPow of n=%d, p=%d was incorrect, got: %d, want: %d", table.n, table.p, digPowResult, table.res)
		}
	}
}
