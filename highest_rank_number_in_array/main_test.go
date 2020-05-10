package highest_rank_number_in_array

import "testing"

func TestHighestRank(t *testing.T) {
	samples := []struct {
		input  []int
		output int
	}{
		{[]int{12, 10, 8, 12, 7, 6, 6, 4, 10, 12}, 12},
		{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}, 12},
		{[]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}, 3},
	}

	for _, sample := range samples {
		result := HighestRank(sample.input)

		if sample.output != result {
			t.Errorf("INCORRECT HighestRank -> GOT: %d, WANT: %d\n", result, sample.output)
		}
	}

}
