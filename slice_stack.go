package gcl

type SliceStack struct {
	size int
	data []interface{}
}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (a *SliceStack) Size() int {
	return a.size
}

func (a *SliceStack) IsEmpty() bool {
	return a.size == 0
}

func (a *SliceStack) Push(value interface{}) {
	a.data = append(a.data, value)
	a.size++
}

func (a *SliceStack) Pop() interface{} {
	a.checkNonEmpty()
	a.size--
	v := a.data[a.size]
	a.data = a.data[:a.size]
	return v
}

func (a *SliceStack) Top() interface{} {
	a.checkNonEmpty()
	return a.data[a.size-1]
}

func (a *SliceStack) Reset() {
	a.data = make([]interface{}, 0)
	a.size = 0
}

func (a *SliceStack) checkNonEmpty() {
	if a.IsEmpty() {
		panic("[SliceStack]: operation on empty stack")
	}
}
