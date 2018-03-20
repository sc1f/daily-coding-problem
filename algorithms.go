package main

import (
	"fmt"
	"github.com/sc1f/algorithms/algorithms"
)

// TODO: write a cleaner runner

func main() {
	sample := make([]int, 20)
	for i := 0; i < 20; i++ {
		sample[i] = i
	}
	fmt.Printf("%d: %v\n", 1, algorithms.ReservoirSampler(sample, 1))
	fmt.Printf("%d: %v\n", 3, algorithms.ReservoirSampler(sample, 3))
	fmt.Printf("%d: %v\n", 5, algorithms.ReservoirSampler(sample, 5))
	fmt.Printf("%d: %v\n", 10, algorithms.ReservoirSampler(sample, 10))
	fmt.Printf("%v\n", algorithms.Shuffle(sample))
	fmt.Println(algorithms.Staircase(12))
}
