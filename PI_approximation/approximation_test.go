package PI_approximation

import "testing"

func TestIterPi(t *testing.T) {
	tables := []struct {
		epsilon          float64
		iterationCounter int
		approximated     string
	}{
		{0.1, 10, "3.0418396189"},
		{0.01, 100, "3.1315929036"},
		{0.001, 1000, "3.1405926538"},
		{0.0001, 10000, "3.1414926536"},
		{0.00001, 100001, "3.1416026535"},
		{0.000001, 1000001, "3.1415936536"},
		{0.05, 20, "3.0916238067"},
	}

	for _, table := range tables {
		iterationResult, approximatedResult := IterPi(table.epsilon)
		if (iterationResult != table.iterationCounter) && (approximatedResult != table.approximated) {
			t.Errorf("IterPi of epsilon=%f, was INCORRECT, GOT: (%d, %s), WANT: (%d, %s)", table.epsilon, iterationResult, approximatedResult, table.iterationCounter, table.approximated)
		}
	}
}
