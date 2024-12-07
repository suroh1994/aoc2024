package main

import "testing"

func TestMyMath(t *testing.T) {
	tt := []struct {
		input  int
		result float64
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{999, 3},
	}

	for _, test := range tt {
		if digitsInNum(test.input) != test.result {
			t.Errorf("%v: got %v; want %v", test, digitsInNum(test.input), test.result)
		}
	}
}
