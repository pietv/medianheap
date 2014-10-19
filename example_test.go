package medianheap_test

import (
	"fmt"

	"github.com/pietv/medianheap"
)

func ExampleIntMedianHeap() {
	h := medianheap.New()
	h.Add(-1)
	h.Add(0)
	h.Add(1)
	fmt.Println(h.Median())
	// Output: 0
}
