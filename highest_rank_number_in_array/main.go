package highest_rank_number_in_array

func countingNums(nums []int) map[int]int {
	result := make(map[int]int)

	for _, v := range nums {
		result[v]++
	}

	return result
}

// HighestRank - returns the number which is most frequent in given slice
func HighestRank(nums []int) int {
	temp := countingNums(nums)
	indexOflast := len(nums) - 1

	return temp[nums[indexOflast]]
}
