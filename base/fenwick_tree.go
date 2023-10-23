package base

import "golang.org/x/exp/constraints"

// 树状数组
type FenwickTree[T constraints.Integer] struct {
	c      []T
	sum    T
	values []T
}

func NewFenwickTree[T constraints.Integer]() *FenwickTree[T] {
	return &FenwickTree[T]{}
}
func (f *FenwickTree[T]) Init(values []T) {
	f.values = values
	f.c = make([]T, len(values)+1)
	for i := 1; i <= len(f.values); i++ {
		f.sum += f.values[i]
		f.c[i] += f.values[i]
		if i+f.lowBit(i) <= len(f.values) {
			f.c[i+f.lowBit(i)] += f.c[i]
		}
	}
}

func (f *FenwickTree[T]) Update(i int, value T) {
	for i <= len(f.values) {
		f.sum += value
		f.c[i] += value
		i += f.lowBit(i)
	}
}

func (f *FenwickTree[T]) Query(i int) (res T) {
	for i > 0 {
		res += f.c[i]
		i -= f.lowBit(i)
	}
	return res
}

func (f *FenwickTree[T]) lowBit(i int) int {
	return i & (-i)
}
