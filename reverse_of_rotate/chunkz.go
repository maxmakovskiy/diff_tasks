package main

import (
	"fmt"
	"math"
)

func isEven(s string) bool {
	if (len(s) % 2) == 0 {
		return true
	} else {
		return false
	}
}

func makeEven(s string) (res string) {
	res = s[:len(s)-1]
	return
}

func rearrangeEven(s string, l int) (res string) {
	if len(s)%l == 0 {
		res = s
	} else {
		// 8 % 3 = 2
		// [0:((8-1) + 1)]
		t := len(s)
		res = s[:(t - (t % l) + 1)]
	}

	return
}

func rearrange(s string, lOfOne int) (rearranged string) {
	if isEven(s) {
		rearranged = rearrangeEven(s, lOfOne)
	} else {
		temp := makeEven(s)
		// fmt.Printf("makeEven(%s) = %s -> len(%s) = %d \n", s, temp, temp, len(temp))
		rearranged = rearrangeEven(temp, lOfOne)
	}

	return
}

func getChunkz(s string, lOfOne int) []string {
	rearrangedString := rearrange(s, lOfOne)
	// fmt.Printf("rearrangedString = %s -> len = %d\n", rearrangedString, len(rearrangedString))

	numOfChunkz := len(rearrangedString) / lOfOne
	// fmt.Printf("number of chunkz -> %d\n", numOfChunkz)

	chunkz := make([]string, numOfChunkz)

	var counter, iterator int
	var chunk string
	for pos, char := range rearrangedString {
		if counter < lOfOne {
			chunk += string(char)
			counter++
			if pos != len(rearrangedString)-1 {
				continue
			}
		}
		chunkz[iterator] = chunk
		counter = 1
		chunk = string(char)
		iterator++
		if iterator == numOfChunkz {
			break
		}
	}

	return chunkz
}

func rotateToLeftByOne(num string) (res string) {
	res = num[1:]
	res += string(num[0])
	return
}

func isSumOfCubesDivisibleByTwo(num string) bool {
	digits := make([]int, len(num))

	for pos, digit := range num {
		digits[pos] = int(digit)
	}

	var temp float64

	for _, digit := range digits {
		temp += math.Pow(float64(digit), 3)
	}

	if (int(temp) % 2) == 0 {
		return true
	}

	return false
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Revrot(s string, n int) (res string) {
	chunkz := getChunkz(s, n)

	for _, v := range chunkz {
		if isSumOfCubesDivisibleByTwo(v) {
			res += reverse(v)
		} else {
			res += rotateToLeftByOne(v)
		}
	}

	return
}

func main() {
	/*
		src1 := "123456987654"
		lOfChunk1 := 6
		fmt.Printf("SOURCE: %s\nLEN of SOURCE: %d\nLEN of ONE: %d\nRESULT: %v\n\n", src1, len(src1), lOfChunk1, getChunkz(src1, lOfChunk1))

		src2 := "66443875"
		lOfChunk2 := 4
		fmt.Printf("SOURCE: %s\nLEN of SOURCE: %d\nLEN of ONE: %d\nRESULT: %v\n\n", src2, len(src2), lOfChunk2, getChunkz(src2, lOfChunk2))

		src3 := "66443875"
		lOfChunk3 := 8
		fmt.Printf("SOURCE: %s\nLEN of SOURCE: %d\nLEN of ONE: %d\nRESULT: %v\n\n", src3, len(src3), lOfChunk3, getChunkz(src3, lOfChunk3))

		src4 := "563000655734469485"
		lOfChunk4 := 2
		fmt.Printf("SOURCE: %s\nLEN of SOURCE: %d\nLEN of ONE: %d\nRESULT: %v\n\n", src4, len(src4), lOfChunk4, getChunkz(src4, lOfChunk4))
	*/

	src1 := "123456987654"
	lOfChunk1 := 6
	fmt.Printf("SOURCE: %s\n --> RESULT: %v\n\n", src1, Revrot(src1, lOfChunk1))

	src2 := "66443875"
	lOfChunk2 := 4
	fmt.Printf("SOURCE: %s\n --> RESULT: %v\n\n", src2, Revrot(src2, lOfChunk2))

	src3 := "123456779"
	lOfChunk3 := 8
	fmt.Printf("SOURCE: %s\n --> RESULT: %v\n\n", src3, Revrot(src3, lOfChunk3))

	// TODO: REVIEW -> WRONG RESULT
	src4 := "563000655734469485"
	lOfChunk4 := 4
	fmt.Printf("SOURCE: %s\n --> RESULT: %v\n\n", src4, Revrot(src4, lOfChunk4))

}
