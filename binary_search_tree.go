package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

func (s *SearchTreeData) Insert(data int) *SearchTreeData {
	if s == nil {
		return Bst(data)
	}
	if data <= s.data {
		s.left = s.left.Insert(data)
	} else {
		s.right = s.right.Insert(data)
	}
	return s
}

func (s *SearchTreeData) MapString(m func(int) string) []string {
	if s == nil {
		return []string{}
	}
	return append(append(s.left.MapString(m), m(s.data)), s.right.MapString(m)...)
}

func (s *SearchTreeData) MapInt(m func(int) int) []int {
	if s == nil {
		return []int{}
	}
	return append(append(s.left.MapInt(m), m(s.data)), s.right.MapInt(m)...)
}
