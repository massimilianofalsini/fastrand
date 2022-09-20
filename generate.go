package fastrand

import (
	"fmt"
)

func Generate(entropy int) int {
	if entropy == 0 {
		entropy = 1000
	}
	results := make(chan int, entropy)
	for i := 0; i < entropy; i++ {
		go func() {
			var num int
			stringNum := fmt.Sprintf("%d\n", &num)
			if _, err := fmt.Sscanf(stringNum, "%d\n", &num); err != nil {
				results <- 0 // send 10 if error
			}
			results <- num % 11
		}()
	}
	var out []int
	for i := 0; i < entropy; i++ {
		out = append(out, <-results)
	}
	return out[entropy-1]
}
