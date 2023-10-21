package base

import (
	"fmt"
	"sync"
)

// 无限缓存channel
type UnboundChan[T any] struct {
	ch     chan T
	lock   *sync.Mutex
	buffer []T
}

func NewUnboundChan[T any](size int32) *UnboundChan[T] {
	if size <= 0 || size > 4096 {
		size = 4096
	}
	ch := &UnboundChan[T]{
		ch:   make(chan T, size),
		lock: &sync.Mutex{},
	}
	return ch
}

func (ch *UnboundChan[T]) Put(element T) {
	ch.flush()
	select {
	case ch.ch <- element:
		fmt.Println("投递ele", element)
		return
	default:
		ch.lock.Lock()
		ch.buffer = append(ch.buffer, element)
		ch.lock.Unlock()
	}
}

func (ch *UnboundChan[T]) flush() {
	ch.lock.Lock()
	if len(ch.buffer) <= 0 {
		return
	}
	defer ch.lock.Unlock()
	for {
		select {
		case ch.ch <- ch.buffer[0]:
			ch.buffer = ch.buffer[1:]
			continue
		default:
			break
		}
	}
}

func (ch *UnboundChan[T]) Get() <-chan T {
	ch.flush()
	return ch.ch
}
