package apply

import (
	"container/list"
	"fmt"
	"time"
)

const (
	interval         = 2500
	TIME_NEAR_SHIFT  = 8
	TIME_NEAR        = 1 << TIME_NEAR_SHIFT
	TIME_LEVEL_SHIFT = 6
	TIME_LEVEL       = 1 << TIME_LEVEL_SHIFT
	TIME_NEAR_MASK   = TIME_NEAR - 1
	TIME_LEVEL_MASK  = TIME_LEVEL - 1
)

type TimeWheelHandle func()

// 时间轮节点
type TimerWheelNode struct {
	expire int64
	handle TimeWheelHandle
}

type TimeWheel struct {
	tick uint32
	near [TIME_NEAR]*list.List
	t    [4][TIME_LEVEL]*list.List
	quit bool
}

func NewTimeWheel() *TimeWheel {
	t := &TimeWheel{}
	for i := range t.near {
		t.near[i] = list.New()
	}
	for i, v := range t.t {
		for j := range v {
			t.t[i][j] = list.New()
		}
	}
	go t.Run()
	return t
}
func (t *TimeWheel) execute() {
	idx := t.tick & TIME_NEAR_MASK
	t.dispatch(t.near[idx])
}

func (t *TimeWheel) AddNode(duration time.Duration, handle TimeWheelHandle) {
	tick := int64(t.tick) + duration.Milliseconds()/int64(10)
	t.addNode(&TimerWheelNode{expire: tick, handle: handle})
}

func (t *TimeWheel) addNode(node *TimerWheelNode) {
	if node.expire|TIME_NEAR_MASK == int64(t.tick)|TIME_NEAR_MASK {
		t.near[node.expire&TIME_NEAR_MASK].PushBack(node)
	}
	mask := TIME_NEAR << TIME_LEVEL_SHIFT
	var i int
	for i = 0; i < 3; i++ {
		if node.expire|int64(mask-1) == int64(t.tick)|int64(mask-1) {
			break
		}
		mask <<= TIME_LEVEL_SHIFT
	}
	t.t[i][node.expire>>(TIME_NEAR_SHIFT+i*TIME_LEVEL_SHIFT)&TIME_LEVEL_MASK].PushBack(node)
}
func (t *TimeWheel) timerShift() {
	mask := TIME_NEAR
	t.tick++
	ct := t.tick
	if ct == 0 {
		t.move(3, 0)
	} else {
		var i int
		it := ct >> TIME_NEAR_SHIFT
		for int(ct)&(mask-1) == 0 {
			idx := int(it) & (mask - 1)
			if idx != 0 {
				t.move(i, idx)
				break
			}
			it >>= TIME_LEVEL_SHIFT
			mask <<= TIME_LEVEL_SHIFT
			i++
		}
	}
}

func (t *TimeWheel) move(level, idx int) {
	e := t.t[level][idx].Front()
	for e != nil {
		t.addNode(e.Value.(*TimerWheelNode))
		e = e.Prev()
	}
}
func (t *TimeWheel) Update() {
	t.update()
}
func (t *TimeWheel) update() {
	t.execute()
	t.timerShift()
	t.execute()
}

func (t *TimeWheel) dispatch(l *list.List) {
	e := l.Front()
	for e != nil {
		l.Remove(e)
		(e.Value.(*TimerWheelNode)).handle()
		e = e.Prev()
	}
}

func (t *TimeWheel) Run() {
	fmt.Println("timeWheel run start")
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		if t.quit {
			break
		}

		t.Update()
		//now := time.Now()
		<-ticker.C
		//fmt.Println(t.tick, time.Since(now).Microseconds())
	}
}

func (t *TimeWheel) Close() {
	fmt.Println("timeWheel close")
	t.quit = true
}
