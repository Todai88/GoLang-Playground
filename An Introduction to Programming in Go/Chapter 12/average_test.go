package main

import (
	"testing"
)

type testpair struct {
	values  []float64
	average float64
	min     float64
	max     float64
}

func max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	temp := xs[0]
	for _, value := range xs {
		if value > temp {
			temp = value
		}
	}
	return temp
}

func min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	temp := xs[0]
	for _, value := range xs {
		if value < temp {
			temp = value
		}
	}
	return temp
}

func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
	{[]float64{}, 0, 0, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		v := max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		v := min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
