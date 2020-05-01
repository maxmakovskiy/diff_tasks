package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negatime number: %d", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	counter := 0

	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	for counter < 10 {
		if (z*z - x) == 0 {
			return z, nil
		} else {
			fmt.Println(z)
			z -= (z*z - x) / (2 * z)
		}

		counter += 1
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
