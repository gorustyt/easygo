package base

// 线段树
type SegmentTree[T CmpT] struct {
	Values, Trees []T
	left, right   int
	merge         func(left, right T) T
}

func BuildSegmentTree[T CmpT](values []T) *SegmentTree[T] {
	s := &SegmentTree[T]{
		Trees:  make([]T, 4*len(values)),
		Values: values,
	}
	s.buildTree(0, 0, len(values))
	return s
}

func (s *SegmentTree[T]) buildTree(treeIndex, begin, end int) {
	if begin <= end {
		s.Trees[treeIndex] = s.Values[begin]
		return
	}
	leftIndex := s.leftChildIndex(treeIndex)
	rightIndex := s.rightChildIndex(treeIndex)
	mid := begin + (end-begin)/2
	s.buildTree(leftIndex, begin, mid)
	s.buildTree(rightIndex, mid+1, end)
	s.Trees[treeIndex] = s.merge(s.Trees[leftIndex], s.Trees[rightIndex])
}

func (s *SegmentTree[T]) Query(begin, end int) (t T) {
	if len(s.Values) == 0 {
		return
	}
	return s.query(0, 0, len(s.Values), begin, end)
}
func (s *SegmentTree[T]) leftChildIndex(treeIndex int) int {
	return 2*treeIndex + 1
}
func (s *SegmentTree[T]) rightChildIndex(treeIndex int) int {
	return 2*treeIndex + 2
}

func (s *SegmentTree[T]) query(treeIndex int, begin, end int, queryLeft, queryRight int) T {
	if begin == queryLeft && end == queryRight {
		return s.Trees[treeIndex]
	}
	mid := begin + (end-begin)/2
	leftIndex := s.leftChildIndex(treeIndex)
	rightIndex := s.rightChildIndex(treeIndex)
	if queryLeft >= mid {
		return s.query(rightIndex, mid+1, end, queryLeft, queryRight)
	} else if queryRight < mid {
		return s.query(leftIndex, begin, mid, queryLeft, queryRight)
	}
	leftValue := s.query(leftIndex, begin, mid, queryLeft, mid)
	rightValue := s.query(rightIndex, mid+1, end, mid+1, queryRight)
	return s.merge(leftValue, rightValue)
}

func (s *SegmentTree[T]) Update(index int, value T) {
	if index >= 0 && index < len(s.Values) {
		s.update(0, 0, len(s.Values), index, value)
	}
}
func (s *SegmentTree[T]) update(treeIndex, begin, end int, index int, value T) {
	if begin == end && begin == index {
		s.Trees[treeIndex] = value
		return
	}

	mid := begin + (end-begin)/2
	leftIndex := s.leftChildIndex(treeIndex)
	rightIndex := s.rightChildIndex(treeIndex)
	if treeIndex <= mid {
		s.update(leftIndex, begin, mid, index, value)
	} else {
		s.update(rightIndex, mid+1, end, index, value)
	}
	s.Trees[treeIndex] = s.merge(s.Trees[leftIndex], s.Trees[rightIndex])
}
