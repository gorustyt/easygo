package base

import (
	"math"
	"math/rand"
)

type SortSet[T SkipElement] struct {
	elements map[string]T
	*SkipList[T]
}

func NewSortSet[T SkipElement]() *SortSet[T] {
	return &SortSet[T]{
		SkipList: NewSkipList[T](),
		elements: map[string]T{}}
}

func (s *SortSet[T]) Add(value T) (exist bool) {
	v, ok := s.elements[value.Key()]
	if ok {
		s.SkipList.Remove(v)
		exist = true
	}
	s.elements[value.Key()] = value
	s.SkipList.Add(value)
	return
}
func (s *SortSet[T]) Get(key string) (value T, ok bool) {
	value, ok = s.elements[key]
	return
}

func (s *SortSet[T]) Remove(key string) (exist bool) {
	v, ok := s.elements[key]
	if ok {
		s.SkipList.Remove(v)
		exist = true
	}
	return
}
func (s *SortSet[T]) GetElementRank(rank int) (value T) {
	return s.SkipList.GetElementByRank(rank)
}
func (s *SortSet[T]) GetRank(key string) int {
	v, ok := s.elements[key]
	if !ok {
		return -1
	}
	return s.SkipList.GetRank(v)
}

func (s *SortSet[T]) DeleteByScore() {

}
func (s *SortSet[T]) RangeByScore() {

}

func (s *SortSet[T]) DeleteByRank() {

}

func (s *SortSet[T]) Card() int64 {
	return int64(s.count)
}

func (s *SortSet[T]) Count(key string, min, max SkipElement) {

}

func (s *SortSet[T]) Rem(keys ...string) {
	for _, v := range keys {
		e, ok := s.elements[v]
		if !ok {
			continue
		}
		s.SkipList.Remove(e)
	}
}

const (
	MaxLevel = 32
)

type SkipElement interface {
	Less(SkipElement) bool
	Key() string //唯一key
}

type SkipList[T SkipElement] struct {
	root     *skipListNode[T]
	count    int
	curLevel int //有效层数
	ranks    []int
}

func NewSkipList[T SkipElement]() *SkipList[T] {
	return &SkipList[T]{curLevel: 1, ranks: make([]int, MaxLevel), root: &skipListNode[T]{levels: make([]*skipListNode[T], MaxLevel)}}
}

func (s *SkipList[T]) Add(value T) {
	r := s.root
	var updates []*skipListNode[T]
	for i := s.curLevel - 1; i >= 0; i-- {
		if i == s.curLevel-1 {
			s.ranks[i] = 0
		} else {
			s.ranks[i] = s.ranks[i+1]
		}
		for r.levels[i] != nil && r.levels[i].value.Less(value) {
			s.ranks[i] += r.levels[i].span
			r = r.levels[i]
		}
		updates = append(updates, r)
	}
	level := s.randomLevel()
	newNode := &skipListNode[T]{value: value, levels: make([]*skipListNode[T], MaxLevel)}
	minLevel := int(math.Min(float64(level), float64(s.curLevel)))

	if minLevel == 0 {
		minLevel = 1
	}
	if minLevel != s.curLevel {
		for i := minLevel; i < s.curLevel; i++ {
			s.root.levels[i] = newNode
			s.ranks[i] = s.count + 1
		}
		s.curLevel = level
	}

	for i := 0; i < minLevel; i++ {
		for _, v := range updates {
			newNode.levels[i] = v.levels[i]
			v.levels[i] = newNode
			newNode.span = s.ranks[i] + v.levels[i].span - s.ranks[0]
			v.levels[i].span = s.ranks[0] - s.ranks[i] + 1
		}
	}
	s.count++
}

func (s *SkipList[T]) randomLevel() (level int) {
	level++
	for rand.Float64() <= 0.25 {
		level++
	}
	return level
}

func (s *SkipList[T]) find(value T) *skipListNode[T] {
	r := s.root
	for i := s.curLevel - 1; i >= 0; i-- {
		for r.levels[i] != nil && r.levels[i].value.Less(value) && !(r.levels[i].value.Key() == value.Key()) {
			r = r.levels[i]
		}
		if r.levels[i] != nil && r.levels[i].value.Key() == value.Key() {
			return r.levels[i]
		}
	}
	return nil
}

func (s *SkipList[T]) Remove(value T) {
	r := s.root
	hasExist := false
	var updates []*skipListNode[T]
	for i := s.curLevel - 1; i >= 0; i-- {
		for r.levels[i] != nil && r.levels[i].value.Less(value) && !(r.levels[i].value.Key() == value.Key()) {
			r = r.levels[i]
		}
		if r.levels[i] != nil && r.levels[i].value.Key() == value.Key() {
			hasExist = true
		}
		updates = append(updates, r)
	}
	if !hasExist { //不存在直接return
		return
	}
	for i := s.curLevel - 1; i >= 0; i-- {
		for _, v := range updates {
			if v.levels[i] != nil {
				v.levels[i] = v.levels[i].levels[i]
			}
		}
	}
	for i := s.curLevel - 1; i >= 0; i-- {
		if s.root.levels[i] != nil {
			break

		}
		s.curLevel--
	}
}

type skipListNode[T SkipElement] struct {
	value  T
	levels []*skipListNode[T]
	span   int
}

// ===============================================================================辅助接口=============================================

func (s *SkipList[T]) GetElementByRank(rank int) (value T) {
	if rank > s.count {
		return
	}
	r := s.root
	span := 0
	for i := s.curLevel - 1; i >= 0; i-- {
		for r.levels[i] != nil && span+r.levels[i].span < rank {
			span += r.levels[i].span
			r = r.levels[i]
		}
		if r.levels[i] != nil && span+r.levels[i].span == rank { //找到了
			return r.levels[i].value
		}
	}
	return
}

func (s *SkipList[T]) GetRank(value T) int {
	r := s.root
	span := 1
	for i := s.curLevel - 1; i >= 0; i-- {
		for r.levels[i] != nil && r.levels[i].value.Less(value) && !(r.levels[i].value.Key() == value.Key()) {
			span += r.levels[i].span
			r = r.levels[i]
		}
		if r.levels[i] != nil && r.levels[i].value.Key() == value.Key() {
			return span
		}
	}
	return -1
}
