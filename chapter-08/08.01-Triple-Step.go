// A child is running up a staircase with n steps and can hop either 1 step, 2
// steps, or 3 steps at a time. Implement a method to count how many possible
// ways the child can run up the stairs.

package main

import "fmt"

// This function is time and space optimized for one call.
// For mulitple calls, possible ways for each step can be stored in a lookup
// table for increased time optimization (i.e. trade-off with space).
func WaysToSteps(steps uint) uint {
	var ways1, ways2, ways3 uint = 1, 2, 4
	if steps == 0 || steps == 1 {
		return ways1
	} else if steps == 2 {
		return ways2
	}

	for i := steps; i > 3; i-- {
		ways3, ways2, ways1 = ways3+ways2+ways1, ways3, ways2
	}

	return ways3
}

func main() {
	for s := 1; s < 10; s++ {
		fmt.Printf("Steps = %v, Ways = %v\n", s, WaysToSteps(uint(s)))
	}
}
