package base

import (
	"cmp"
	"math/rand"
)

const (
	MaxLevel = 32
)

type SkipList[T cmp.Ordered] struct {
	root  *skipListNode[T]
	count int64
}

func NewSkipList[T cmp.Ordered]() *SkipList[T] {
	return &SkipList[T]{root: &skipListNode[T]{levels: make([]*skipListNode[T], MaxLevel)}}
}

func (s SkipList[T]) Add(value T) {

}
func (s SkipList[T]) randomLevel() (level int) {
	for rand.Float64() <= 0.25 {
		level++
	}
	return level
}

func (s SkipList[T]) Remove(value T) {

}

type skipListNode[T cmp.Ordered] struct {
	value  T
	levels []*skipListNode[T]
}
