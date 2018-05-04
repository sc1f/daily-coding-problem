package main

import (
	"fmt"
	"github.com/sc1f/algorithms/algorithms"
)

// TODO: write a cleaner runner

func main() {
	sample := make([]int, 20)
	for i := 0; i < 20; i++ {
		sample[i] = i + 1
	}
	fmt.Println(sample)
	fmt.Println(algorithms.TwoSumPointers(sample, 13))
	fmt.Println(algorithms.ThreeSum(sample, 12))
	fmt.Printf("%v\n", algorithms.Shuffle(sample))
	fmt.Println(algorithms.InsertionSort(sample))
	/*
		fmt.Printf("%d: %v\n", 1, algorithms.ReservoirSampler(sample, 1))
		fmt.Printf("%d: %v\n", 3, algorithms.ReservoirSampler(sample, 3))
		fmt.Printf("%d: %v\n", 5, algorithms.ReservoirSampler(sample, 5))
		fmt.Printf("%d: %v\n", 10, algorithms.ReservoirSampler(sample, 10))
		fmt.Println(algorithms.Staircase(12))
		//fmt.Println(algorithms.LongestUniqueSubstring(15, "abcba"))
	*/
}
