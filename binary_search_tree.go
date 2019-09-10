package gcl

type BiSearchTree struct {
	size          int
	compareMethod func(v1, v2 interface{}) int
	root          *biSearchTreeNode
}

func NewBiSearchTree(compareMethod func(v1, v2 interface{}) int) *BiSearchTree {
	return &BiSearchTree{compareMethod: compareMethod}
}

func (b *BiSearchTree) Size() int {
	return b.size
}

func (b *BiSearchTree) IsEmpty() bool {
	return b.size == 0
}

func (b *BiSearchTree) Insert(value interface{}) {
	z := &biSearchTreeNode{value: value}
	var y *biSearchTreeNode
	x := b.root
	for x != nil {
		y = x
		if b.compareMethod(z.value, x.value) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	if y == nil {
		b.root = z
	} else if b.compareMethod(z.value, y.value) < 0 {
		y.left = z
	} else {
		y.right = z
	}

	b.size++
}

func (b *BiSearchTree) Remove(value interface{}) {

}

func (b *BiSearchTree) PreTravel() []interface{} {
	ret := make([]interface{}, 0)
	if !b.IsEmpty() {
		node := b.root
		ret = b.preTravel(ret, node)
	}
	return ret
}

func (b *BiSearchTree) preTravel(ret []interface{}, node *biSearchTreeNode) []interface{} {
	if node != nil {
		ret = append(ret, node.value)
		ret = b.preTravel(ret, node.left)
		ret = b.preTravel(ret, node.right)
	}
	return ret
}

type biSearchTreeNode struct {
	value       interface{}
	left, right *biSearchTreeNode
}
