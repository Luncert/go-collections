package gcl

type SliceList struct {
	size        int
	matchMethod func(v1, v2 interface{}) bool
	data        []interface{}
}

func NewSliceList(matchMethod func(v1, v2 interface{}) bool) *SliceList {
	return &SliceList{
		matchMethod: matchMethod,
		data:        make([]interface{}, 0),
	}
}

func (s *SliceList) Size() int {
	return s.size
}

func (s *SliceList) IsEmpty() bool {
	return s.size == 0
}

func (s *SliceList) Get(index int) interface{} {
	s.checkIndex(index)
	return s.data[index]
}

func (s *SliceList) Index(value interface{}) int {
	for i, v := range s.data {
		if s.matchMethod(value, v) {
			return i
		}
	}
	return -1
}

func (s *SliceList) ForEach(call func(value interface{})) {
	for _, v := range s.data {
		call(v)
	}
}

func (s *SliceList) Append(value interface{}) List {
	s.data = append(s.data, value)
	return s
}

func (s *SliceList) Pop() interface{} {
	s.checkNonEmpty()
	s.size--
	v := s.data[s.size]
	s.data = s.data[:s.size]
	return v
}

func (s *SliceList) Set(index int, value interface{}) interface{} {
	s.checkIndex(index)
	oldValue := s.data[index]
	s.data[index] = value
	return oldValue
}

func (s *SliceList) Remove(index int) interface{} {
	s.checkIndex(index)
	v := s.data[index]
	s.data = append(s.data[:index], s.data[index+1:])
	s.size--
	return v
}

func (s *SliceList) Reset() {
	s.size = 0
	s.data = s.data[:]
}

func (s *SliceList) checkIndex(index int) {
	if index < 0 || index >= s.size {
		panic("[SliceList]: index out of range")
	}
}

func (s *SliceList) checkNonEmpty() {
	if s.size == 0 {
		panic("[SliceList]: operation on empty collection")
	}
}
