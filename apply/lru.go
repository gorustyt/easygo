package apply

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

type Lru interface {
	Len() int
	Keys() []string
	GetOldest() (key string, value any, ok bool)
	Remove(key string) (present bool)
	RemoveOldest() (key string, value any, ok bool)
	Resize(size int) (evicted int)
	PeekOrAdd(key string, value any) (previous any, ok, evicted bool)
	ContainsOrAdd(key string, value any) (ok, evicted bool)
	Peek(key string) (value any, ok bool)
	Contains(key string) bool
	Get(key string) (value any, ok bool)
	Add(key string, value any) (evicted bool)
	Purge()
}

type lru struct {
	dict      map[string]*base.Element
	Size      int64
	List      base.DoubleLinkList
	onEvictCb func(key string, value interface{})
	Cap       int
}

type lruEntry struct {
	key string
	ele interface{}
}

func (entry *lruEntry) GetKey() (key string) {
	return entry.key
}

func (entry *lruEntry) GetValue() (ele interface{}) {
	e := entry
	return e
}

func NewLruEntry(key string, ele interface{}) *lruEntry {
	return &lruEntry{
		key: key,
		ele: ele,
	}
}

func NewLru(cap int, onEvictCb func(key string, value interface{})) Lru {
	if cap < 0 {
		panic(any("cap should has value"))
	}
	return &lru{
		dict:      map[string]*base.Element{},
		List:      base.NewDoubleLinkList(),
		Cap:       cap,
		onEvictCb: onEvictCb,
	}
}

func (l *lru) Remove(key string) (present bool) {
	node, ok := l.dict[key]
	if !ok {
		return false
	}
	l.removeElement(key, node)
	return true
}

func (l *lru) Len() int {
	return int(l.Size)
}
func (l *lru) Keys() []string {
	var keys []string
	for k := range l.dict {
		keys = append(keys, k)
	}
	return keys
}

func (l *lru) GetOldest() (key string, value any, ok bool) {
	ele := l.List.Back()
	if ele == nil {
		return "", ele, false
	}
	en, ok := ele.Value.(*lruEntry)
	if !ok {
		return "", ele, false
	}
	return en.GetKey(), en.GetValue(), true
}

func (l *lru) RemoveOldest() (key string, value any, ok bool) {
	ele := l.List.Back()
	if ele == nil {
		return "", ele, false
	}
	en, ok := ele.Value.(*lruEntry)
	if !ok {
		return "", ele, false
	}
	node, ok := l.dict[en.GetKey()]
	if !ok {
		panic(any(fmt.Sprintf("there expect key:%v has value", en.GetKey())))
	}
	l.removeElement(key, node)
	return en.GetKey(), en.GetValue(), true
}

func (l *lru) Resize(size int) (evicted int) {
	if size <= 0 {
		return
	}
	l.Cap = size
	if l.List.Len() > size {
		length := l.List.Len() - size
		for i := 0; i < length; i++ {
			e := l.List.Back()
			en, ok := e.Value.(*lruEntry)
			if ok {
				l.removeElement(en.key, e)
			}

		}
		return length
	}
	return 0
}

func (l *lru) PeekOrAdd(key string, value any) (previous any, ok, evicted bool) {
	value, ok = l.Peek(key)
	if !ok {
		return
	}
	evicted = l.Add(key, value)
	return value, true, evicted
}

func (l *lru) ContainsOrAdd(key string, value any) (ok, evicted bool) {
	if !l.Contains(key) {
		return true, l.Add(key, value)
	}
	return
}

func (l *lru) Peek(key string) (value any, ok bool) {
	value, ok = l.dict[key]
	return
}

func (l *lru) Contains(key string) bool {
	_, ok := l.dict[key]
	return ok
}

func (l *lru) Get(key string) (value any, ok bool) {
	node, ok := l.dict[key]
	if !ok {
		return nil, false
	}
	l.List.MoveToFront(node)
	return node.Value, true
}

func (l *lru) Add(key string, value any) (evicted bool) {
	node, ok := l.dict[key]
	if ok {
		node.Value = value
		l.List.MoveToFront(node)
	} else {
		node = l.List.PushFront(NewLruEntry(key, value))
		l.dict[key] = node
		l.Size++
	}
	if l.List.Len() > l.Cap {
		l.RemoveOldest()
		return true
	}
	return
}
func (l *lru) Purge() {
	for k, v := range l.dict {
		l.removeElement(k, v)
	}
}

func (l *lru) removeElement(key string, node *base.Element) {
	delete(l.dict, key)
	l.List.Remove(node)
	if l.onEvictCb != nil {
		l.onEvictCb(key, node.Value)
	}
	l.Size--
}
