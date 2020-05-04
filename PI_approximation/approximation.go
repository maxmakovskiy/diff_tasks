package PI_approximation

import (
	"math"
	"strconv"
)

func seriesOfLeibnizAt(n int) (result float64) {
	for i := 0; i < n; i++ {
		result += math.Pow(-1.0, float64(i)) / float64((2*n)+1)
	}

	return
}

func IterPi(epsilon float64) (int, string) {
	var n int
	var iterationCounter int
	for {
		if t := 4 * seriesOfLeibnizAt(n); epsilon > (math.Pi - t) {
			s := strconv.FormatFloat(t, 'f', -1, 64)
			return iterationCounter, s
		}
		n += 4
		iterationCounter++
	}
}
