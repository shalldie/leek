package utils

import (
	"time"
)

type IntervalTimer struct {
	ticker *time.Ticker
	quit   chan bool
}

// 启动定时器（适用于周期性定时器）
func (t *IntervalTimer) run(callback func()) {
	for {
		select {
		case <-t.ticker.C:
			callback()
		case <-t.quit:
			return
		}
	}
}

// 停止定时器
func (t *IntervalTimer) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
	close(t.quit)
}

// 重启定时器
func (t *IntervalTimer) Restart(duration time.Duration) {
	t.Stop()
	if t.ticker != nil {
		t.ticker = time.NewTicker(duration)
		// go t.run(func() {
		// 	fmt.Println("Periodic task triggered")
		// })
	}
}

// 创建一个周期性定时器
func NewIntervalTimer(duration time.Duration, callback func()) *IntervalTimer {
	t := &IntervalTimer{
		quit: make(chan bool),
	}
	t.ticker = time.NewTicker(duration)
	go t.run(callback)
	return t
}
