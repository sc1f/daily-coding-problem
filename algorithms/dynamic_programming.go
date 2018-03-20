package algorithms

func Staircase(n int) int {
	//There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
	//Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.
	first := 1
	second := 1
	if n == 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		third := first + second
		first, second = second, third
	}
	return second
}

func StaircasePaths(n int) []int {
	//There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
	//Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.
	return make([]int, 3)
}

func SumNonAdjacent(nums []int) int {
	// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
	// For example, [2, 4, 6, 8] should return 12, since we pick 4 and 8. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
	return 0
}

func DecodeMessage(message string) int {
	// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
	// For example, the message '111' would give 3, since it could be decoded as 'aaa, 'ka', and 'ak'.
	return 0
}
