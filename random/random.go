package random

import "math/rand/v2"

// RandRange returns a pseudo-random integer in the range [min, max].
// It uses the automatically-seeded, concurrency-safe `math/rand/v2` package.
func RandRange(min, max int) int {
	// rand.IntN(n) returns a random number in [0, n).
	// To get a number in [min, max], we need a range of (max - min + 1).
	// We generate a random number in [0, max - min + 1) and add min to it.
	return min + rand.IntN(max-min+1)
}
