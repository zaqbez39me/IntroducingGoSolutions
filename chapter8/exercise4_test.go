package main

import "testing"

type testpair struct {
	values   []float64
	expected float64
}

func TestMin(t *testing.T) {
	testValues := []testpair{
		{[]float64{1, 2}, 1},
		{[]float64{1, 1, 1, 1, 1, 1}, 1},
		{[]float64{-1, 1}, -1},
		{[]float64{}, 0},
	}
	for _, pair := range testValues {
		v := Min(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	testValues := []testpair{
		{[]float64{1, 2}, 2},
		{[]float64{1, 1, 1, 1, 1, 1}, 1},
		{[]float64{-1, 1}, 1},
		{[]float64{}, 0},
	}
	for _, pair := range testValues {
		v := Max(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
