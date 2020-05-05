package PI_approximation

import (
	"math"
	"strconv"
)

func seriesOfLeibnizAt(n int) (result float64) {
	for i := 0; i <= n; i++ {
		temp := math.Pow(-1.0, float64(i)) / float64((2*i)+1)
		result += temp
	}

	return
}

func IterPi(epsilon float64) (int, string) {
	var n int = 1
	var iterationCounter int
	for {
		t := 4.0 * seriesOfLeibnizAt(n)
		if epsilon > math.Abs(math.Pi-t) {
			s := strconv.FormatFloat(t, 'f', 10, 64)
			//			s := fmt.Sprintf("%.10f", t)
			return iterationCounter, s
		}
		n++
		iterationCounter++
	}
}
