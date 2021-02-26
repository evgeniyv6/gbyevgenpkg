package slicegen

import "math/rand"

// GenerSlice generates slice with numbers > 0 and < 100
func GenerSlice(num int) []int {
	randNum := make([]int, num)
	for i := 0; i < num; i++ {
		randNum[i] = rand.Intn(100)
	}
	return randNum
}

// ReverseSlice the given slice
func ReverseSlice(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
