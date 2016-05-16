package main

import (
	"bytes"
	"math"
	"os"
	"fmt"
	"strconv"
)

// Parse the command line arguments
func main() {
	argsWithoutProg := os.Args[1:]
	var vals []float64
	for _, s := range argsWithoutProg {
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			vals = append(vals, f)
		}
	}

	fmt.Println(Signals(vals))
}

var ticks = []rune("▁▂▃▄▅▆▇█")

// Generate the normalised signal chart
func Signals(nums []float64) string {
	if len(nums) == 0 {
		return ""
	}
	min, max := bounds(nums)
	var sparks bytes.Buffer
	if math.Abs(max - min) < 0.0000001 {
		for c := 0; c < len(nums); c++ {
			sparks.WriteRune(ticks[0])
		}
	} else {
		scale := 7.0 / (max - min)
		for _, n := range nums {
			tick := int((n - min) * scale + 0.5) // always positive
			sparks.WriteRune(ticks[tick])
		}
	}

	return sparks.String()
}

// Compute the min and max bounds of the dataset
func bounds(nums []float64) (float64, float64) {
	min, max := math.MaxFloat64, math.SmallestNonzeroFloat64
	for _, n := range nums {
		if (n < min) { min = n }
		if (n > max) { max = n }
	}
	return min, max
}