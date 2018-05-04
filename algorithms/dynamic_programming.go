package algorithms

import "math"

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialHelper(n-1)
}

func Staircase(n int) int {
	//There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
	//Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.
	first := 1
	second := 1
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
	inclusive, exclusive := 0, 0
	for _, num := range nums {
		new_exclusive := math.MaxInt32(inclusive, exclusive)
		inclusive = exclusive + num
		exclusive = new_exclusive
	}
	return math.MaxInt32(inclusive, exclusive)
}

func DecodeMessage(message string) int {
	// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
	// For example, the message '111' would give 3, since it could be decoded as 'aaa, 'ka', and 'ak'.
	return 0
}

func MaximumProfit(prices []int64) int64 {
	/* Given a array of numbers representing the stock prices of a company in chronological order, write a function that calculates the maximum profit you could have made from buying and selling that stock once. You must buy before you can sell it.
	For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at 5 dollars and sell it at 10 dollars. */
	max_profit, min_price := prices[1]-prices[0], prices[0]
	for int i = 1; i < len(prices); i++ {
		current_price = prices[i]
		min_price = math.MinInt64(min_price, current_price)
		current_profit := current_price - min_price
		max_profit = math.MaxInt64(max_profit, current_profit)
	}
	return max_profit
}
