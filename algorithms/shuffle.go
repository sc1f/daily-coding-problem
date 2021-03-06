package algorithms

import "math/rand"

func Shuffle(nums []int) []int {
	// Given an array, shuffle it in one o(n) pass.
	n := len(nums) - 1
	for ; n > 0; n-- {
		r := rand.Intn(len(nums))
		nums[n], nums[r] = nums[r], nums[n]
	}
	return nums
}
