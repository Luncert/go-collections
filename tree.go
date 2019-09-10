package gcl

type Tree interface {
	Size() int
	IsEmpty() bool
	Insert(value interface{})
	Reset()
}
