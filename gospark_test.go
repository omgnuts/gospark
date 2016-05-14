package main

import (
	"testing"
)

func expect(t *testing.T, in []float64, out string, result string) {
	if result != out {
		t.Errorf("Error (%v) = %v, want %v", in, result, out)
	}
}

// holman tests https://github.com/holman/spark
func TestHolman(t *testing.T) {
	//it_graphs_argv_data
	in := []float64{1, 5, 22, 13, 5}
	out := "▁▂█▅▂"
	expect(t, in, out, Signals(in))

	//it_charts_pipe_data
	in = []float64{0, 30, 55, 80, 33, 150}
	out = "▁▂▃▅▂█"
	expect(t, in, out, Signals(in))

	//it_handles_decimals
	in = []float64{5.5, 20}
	out = "▁█"
	expect(t, in, out, Signals(in))

	//it_charts_100_lt_300
	in = []float64{1, 2, 3, 4, 100, 5, 10, 20, 50, 300}
	out = "▁▁▁▁▃▁▁▁▂█"
	expect(t, in, out, Signals(in))

	//it_charts_50_lt_100
	in = []float64{1, 50, 100}
	out = "▁▄█"
	expect(t, in, out, Signals(in))

	//it_charts_4_lt_8
	in = []float64{2,4,8}
	out = "▁▃█"
	expect(t, in, out, Signals(in))

	//it_charts_no_tier_0
	in = []float64{1,2,3,4,5}
	out = "▁▃▅▇█"
	expect(t, in, out, Signals(in))
}

func TestGoSpark(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	out := "▁▂▃▄▅▆▇█"
	expect(t, in, out, Signals(in))

	in = []float64{1, 0, 0, 1}
	out = "█▁▁█"
	expect(t, in, out, Signals(in))

	in = []float64{67, 71, 77, 85, 95, 104, 106, 105, 100, 89, 76, 66}
	out = "▁▂▃▄▆███▇▅▃▁"
	expect(t, in, out, Signals(in))
}

func TestEmptyValues(t *testing.T) {
	in := []float64{}
	out := ""
	expect(t, in, out, Signals(in))
}

// Minimal ticks if they are of the same value
func TestSameValues(t *testing.T) {
	in := []float64{1}
	out := "▁"
	expect(t, in, out, Signals(in))

	in = []float64{1, 1, 1, 1}
	out = "▁▁▁▁"
	expect(t, in, out, Signals(in))
}
