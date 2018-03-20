package algorithms

import (
	"math/rand"
)

func ReservoirSampler(stream []int, k int) []int {
	// Given a stream of elements too large to store in memory, pick k random elements from the stream with uniform probability.
	n := len(stream)
	reservoir := make([]int, k)
	// get the first k - 1 items from the stream
	i := 0
	for ; i < k; i++ {
		reservoir[i] = stream[i]
	}
	// from stream[k] to stream[n], pick a random element and replace element in reservoir with element from stream
	for ; i < n; i++ {
		random_index := rand.Intn(n)
		if random_index < k {
			reservoir[random_index] = stream[i]
		}
	}
	return reservoir
}
