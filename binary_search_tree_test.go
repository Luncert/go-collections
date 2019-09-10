package gcl

import "fmt"

func biSearchTreeCompareMethod(v1, v2 interface{}) (r int) {
	a := v1.(int)
	b := v2.(int)
	if a > b {
		r = 1
	} else if a < b {
		r = -1
	} else {
		r = 0
	}
	return
}

func ExampleBiSearchTree_Insert() {
	t := NewBiSearchTree(biSearchTreeCompareMethod)
	t.Insert(1)
	t.Insert(2)
	t.Insert(-1)
	t.Insert(2)
	t.Insert(-2)
	t.Insert(-3)
	fmt.Print(t.PreTravel())

	// Output: [1 -1 -2 -3 2 2]
}
