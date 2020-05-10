package highest_rank_number_in_array

func countingNums(nums []int) map[int]int {
	result := make(map[int]int)

	for _, v := range nums {
		result[v]++
	}

	return result
}

func biggerValue(src map[int]int) (key int) {
	var val int
	for k, v := range src {
		if v > val {
			val = v
			key = k
		}
	}
	return
}

// HighestRank - returns the number which is most frequent in given slice
func HighestRank(nums []int) int {
	temp := countingNums(nums)
	val := biggerValue(temp)
	return val
}
