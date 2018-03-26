package algorithms

import "fmt"

func TwoSum(arr []int, target int) []int {
	// use if list is not sorted: hashtable approach
	seen := make(map[int]int)
	for _, num := range arr {
		part_ans := target - num
		if _, ok := seen[num]; ok {
			return []int{part_ans, num}
		}
		seen[part_ans] = num
	}
	return []int{}
}

func TwoSumPointers(arr []int, target int) []int {
	// use if list is sorted: binary search
	start_index, end_index := 0, len(arr)-1
	for start_index != end_index {
		start := arr[start_index]
		end := arr[end_index]
		sum := start + end
		fmt.Println(sum)
		if sum == target {
			return []int{start, end}
		} else if sum > target {
			end_index--
		} else if sum < target {
			start_index++
		}
	}
	return []int{}
}

func ThreeSum(arr []int, target int) []int {
	start_index, end_index := 1, len(arr)-1
	for _, elem := range arr {
		for start_index != end_index {
			start := arr[start_index]
			end := arr[end_index]
			sum := start + end + elem
			fmt.Println(sum)
			if sum == target {
				return []int{start, elem, end}
			} else if sum > target {
				end_index--
			} else if sum < target {
				start_index++
			}
		}
	}
	return []int{}
}
