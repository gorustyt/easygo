package base

type Queue interface {
	Poll() (ele interface{})
	Offer(ele interface{})
	Len() int
	Empty() bool
}

type SimpleQueue struct {
	elements []interface{}
}

func NewSimpleQueue() Queue {
	return &SimpleQueue{}
}
func (q *SimpleQueue) Poll() (ele interface{}) {
	if q.Empty() {
		return nil
	}
	ele = q.elements[0]
	q.elements = q.elements[1:]
	return ele
}
func (q *SimpleQueue) Offer(ele interface{}) {
	q.elements = append(q.elements, ele)
}
func (q *SimpleQueue) Len() int {
	return len(q.elements)
}
func (q *SimpleQueue) Empty() bool {
	return q.Len() == 0
}
