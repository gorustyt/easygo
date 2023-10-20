package apply

import "time"

type Task func()

type TimeWheel interface {
}

// 时间轮算法
type timeWheel struct {
	interval int
	num      int
	pos      int
	ticker   *time.Ticker
}

func NewTimeWheel(interval int, num int) TimeWheel {
	t := &timeWheel{
		interval: interval,
		num:      num,
		ticker:   time.NewTicker(time.Duration(interval) * time.Second),
	}
	go t.run()
	return t
}

func (t *timeWheel) AddTask(key string, task Task) {

}

func (t *timeWheel) DelTask(key string) {

}

func (t *timeWheel) run() {

}

func (t *timeWheel) Stop() {

}
