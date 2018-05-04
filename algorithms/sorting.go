package algorithms

func BubbleSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		pos := i

		for ok := (pos > 0 && nums[pos-1] > val); ok; {
			nums[pos] = nums[pos-1]
			pos--
		}

		nums[pos] = val
	}

	return nums
}
