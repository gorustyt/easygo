package apply

import (
	"container/list"
	"fmt"
	"time"
)

const (
	TIME_NEAR_SHIFT  = 8
	TIME_NEAR        = 1 << TIME_NEAR_SHIFT
	TIME_LEVEL_SHIFT = 6
	TIME_LEVEL       = 1 << TIME_LEVEL_SHIFT
	TIME_NEAR_MASK   = TIME_NEAR - 1
	TIME_LEVEL_MASK  = TIME_LEVEL - 1
)

type TimeWheelHandle func(ts time.Time)

// 时间轮节点
type TimerWheelNode struct {
	expire       int64
	handle       TimeWheelHandle
	duration     time.Duration
	expireAt     time.Time
	fireDuration time.Duration
}

type TimeWheel struct {
	tick         int64 //不用考虑溢出的问题
	near         [TIME_NEAR]*list.List
	t            [4][TIME_LEVEL]*list.List
	current      int64     //当前跑了多少tick
	currentPoint time.Time //上一次计算时间
	startTime    time.Time //创建时间
	quit         bool      //是否退出
}

func NewTimeWheel() *TimeWheel {
	t := &TimeWheel{
		currentPoint: time.Now(),
		startTime:    time.Now(),
		current:      time.Now().UnixNano() / (10 * 1e6),
	}
	for i := range t.near {
		t.near[i] = list.New()
	}
	for i, v := range t.t {
		for j := range v {
			t.t[i][j] = list.New()
		}
	}
	return t
}

func (t *TimeWheel) execute(ts time.Time) {
	idx := t.tick & TIME_NEAR_MASK
	l := t.near[idx]
	e := l.Front()

	for e != nil {
		l.Remove(e)
		node := e.Value.(*TimerWheelNode)
		if node.fireDuration > 0 { //循环任务
			node.expireAt = node.expireAt.Add(node.duration)
			node.expire = t.calTick(node.expireAt)
			t.addNode(node)
		}
		node.handle(ts)
		e = e.Prev()
	}

}

func (t *TimeWheel) Schedule(fireDuration, duration time.Duration, handle TimeWheelHandle) {
	node := &TimerWheelNode{
		handle:       handle,
		duration:     duration,
		fireDuration: fireDuration,
		expireAt:     time.Now().Add(fireDuration)}
	node.expire = t.calTick(node.expireAt)
	t.addNode(node)

}

func (t *TimeWheel) calTick(expireAt time.Time) int64 {
	return int64(durationToTick(expireAt.Sub(t.startTime)))
}

func (t *TimeWheel) Add(duration time.Duration, handle TimeWheelHandle) {
	node := &TimerWheelNode{
		handle:   handle,
		duration: duration,
		expireAt: time.Now().Add(duration)}
	node.expire = t.calTick(node.expireAt)
	t.addNode(node)
}

func (t *TimeWheel) addNode(node *TimerWheelNode) {
	if node.expire|TIME_NEAR_MASK == t.tick|TIME_NEAR_MASK {
		t.near[node.expire&TIME_NEAR_MASK].PushBack(node)
	} else {
		mask := int64(TIME_NEAR << TIME_LEVEL_SHIFT)
		var i int
		for i = 0; i < 3; i++ {
			if node.expire|(mask-1) == t.tick|(mask-1) {
				break
			}
			mask <<= TIME_LEVEL_SHIFT
		}
		t.t[i][(node.expire>>(TIME_NEAR_SHIFT+i*TIME_LEVEL_SHIFT))&TIME_LEVEL_MASK].PushBack(node)
	}
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
			idx := int(it) & TIME_LEVEL_MASK
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
	l := t.t[level][idx]
	e := l.Front()
	for e != nil {
		l.Remove(e)
		t.addNode(e.Value.(*TimerWheelNode))
		e = e.Prev()
	}
}
func (t *TimeWheel) Update() {
	now := time.Now()
	if now.Before(t.currentPoint) {
		fmt.Println("find error currentPoint")
		t.currentPoint = now
	} else {
		diff := int(durationToTick(now.Sub(t.currentPoint)))
		if diff == 0 {
			return
		}
		t.currentPoint = t.currentPoint.Add(time.Duration(diff) * 10 * time.Millisecond)
		for i := 0; i < diff; i++ {
			t.update(t.currentPoint)
		}
		t.current += int64(diff)
	}
}
func (t *TimeWheel) update(ts time.Time) {
	t.execute(ts)
	t.timerShift()
	t.execute(ts)
}

// 这里改为手动运行，方便自己使用驱动
func (t *TimeWheel) Run() {
	fmt.Println("timeWheel run start")
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t.Update()
		}
		if t.quit {
			break
		}
	}
}

func (t *TimeWheel) Close() {
	fmt.Println("timeWheel close")
	t.quit = true
}

func durationToTick(d time.Duration) float64 {
	sec := d / time.Second
	nsec := d % time.Second
	return float64(sec)*100 + float64(nsec)/1e7
}
