package base

// 线段树
type SegmentTree[T CmpT] struct {
	Values, Trees, lazy []T
	left, right         int
	merge               func(left, right T) T
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
	return s.query(0, 0, len(s.Values)-1, begin, end)
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
		s.update(0, 0, len(s.Values)-1, index, value)
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

func (s *SegmentTree[T]) UpdateLazy(left, right int, value T) {
	if len(s.Values) > 0 {
		s.updateLazy(0, 0, len(s.Values)-1, left, right, value)
	}
}

func (s *SegmentTree[T]) updateLazy(treeIndex int, left, right int, updateLeft, updateRight int, value T) {
	mid := left + (right-left)/2
	leftIndex := s.leftChildIndex(treeIndex)
	rightIndex := s.rightChildIndex(treeIndex)
	var tmp T
	if s.lazy[treeIndex] != tmp {
		for i := 0; i < right-left+1; i++ {
			s.Trees[treeIndex] = s.merge(s.Trees[treeIndex], s.lazy[treeIndex])
		}
		if left != right {
			s.lazy[leftIndex] = s.merge(s.lazy[treeIndex], s.lazy[leftIndex])
			s.lazy[rightIndex] = s.merge(s.lazy[treeIndex], s.lazy[rightIndex])
		}
		s.lazy[treeIndex] = tmp
	}
	if left > right || updateLeft > updateRight || updateRight < left || updateLeft > right {
		return
	}
	if updateLeft <= left && right <= updateRight {
		for i := 0; i < right-left+1; i++ {
			s.Trees[treeIndex] = s.merge(s.Trees[treeIndex], value)
		}
		if left != right {
			s.lazy[leftIndex] = s.merge(value, s.lazy[leftIndex])
			s.lazy[rightIndex] = s.merge(value, s.lazy[rightIndex])
		}
		return
	}
	s.updateLazy(leftIndex, left, mid, updateLeft, updateRight, value)
	s.updateLazy(right, mid+1, right, updateLeft, updateRight, value)
	s.Trees[treeIndex] = s.merge(s.Trees[leftIndex], s.Trees[rightIndex])
}

func (s *SegmentTree[T]) QueryLazy(left, right int, value T) (res T) {
	if len(s.Values) > 0 {
		return s.queryLazy(0, 0, len(s.Values)-1, left, right, value)
	}
	return res
}

func (s *SegmentTree[T]) queryLazy(treeIndex int, left, right int, queryLeft, queryRight int, value T) (res T) {
	if queryRight < left || left > queryRight {
		return
	}
	mid := left + (right-left)/2
	leftIndex := s.leftChildIndex(treeIndex)
	rightIndex := s.rightChildIndex(treeIndex)
	if s.lazy[treeIndex] != res {
		for i := 0; i < right-left+1; i++ {
			s.Trees[treeIndex] = s.merge(s.Trees[treeIndex], s.lazy[treeIndex])
		}
		if left != right {
			s.lazy[leftIndex] = s.merge(s.lazy[treeIndex], s.lazy[leftIndex])
			s.lazy[rightIndex] = s.merge(s.lazy[treeIndex], s.lazy[rightIndex])
		}
		s.lazy[treeIndex] = res
	}
	if queryLeft <= left && queryRight >= right {
		return s.Trees[treeIndex]
	}
	if queryRight <= mid {
		return s.queryLazy(leftIndex, left, mid, queryLeft, queryRight, value)
	} else if queryLeft > mid {
		return s.queryLazy(rightIndex, mid+1, right, queryLeft, queryRight, value)
	}
	return s.merge(s.queryLazy(leftIndex, left, mid+1, queryLeft, mid+1, value), s.queryLazy(rightIndex, mid+1, right, mid+1, queryRight, value))
}
