package core

import (
	"fmt"
	"time"
)

type Timer struct {
	stopFlag bool
	interval int // millsec
	timer    *time.Ticker
	TimeOut  func()
}

func (t *Timer) Start(interval int) {
	t.stopFlag = false
	t.timer = time.NewTicker(time.Millisecond * time.Duration(interval))

	for { //循环
		if t.stopFlag {
			break
		}
		<-t.timer.C
		if t.TimeOut != nil {
			t.TimeOut()
		}
	}
	t.timer = nil
}

func (t *Timer) Stop() {
	t.stopFlag = true
}

func (t *Timer) Restart(interval int) {
	fmt.Println("restart ", interval)
	//t.Stop()
	//t.Start(interval)
	t.timer.Reset(time.Millisecond * time.Duration(interval))
}
