package heap

import (
	"testing"
)

var minComp = func(i1, i2 int) bool { return i1 < i2 }
var maxComp = func(i1, i2 int) bool { return i1 > i2 }

var tt = []struct {
	name string
	comp func(int, int) bool
	in   []int
}{
	{
		name: "MinHeap",
		comp: minComp,
		in:   []int{6, 3, 8, 2, 1},
	},
	{
		name: "MaxHeap",
		comp: maxComp,
		in:   []int{6, 3, 8, 2, 1},
	},
}

func TestHeap(t *testing.T) {
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			heap := NewHeap(tc.comp, len(tc.in))
			for _, val := range tc.in {
				heap.Set(val)
			}

			out := []int{}
			for range tc.in {
				out = append(out, heap.Get())
			}

			for i := 1; i < len(tc.in); i++ {
				if tc.comp(out[i-1], out[i]) {
					continue
				}
				t.Errorf("Heap order not followed at index %d [val: %d] and %d [val: %d]\n", i-1, out[i-1], i, out[i])
				t.FailNow()
			}
		})
	}
}
