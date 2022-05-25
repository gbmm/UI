package example

import (
	"UI/gogui/core"
	"fmt"
	"strconv"
)

type UIQuicksort struct {
	Widget *core.Wnd
	Timer  *core.Timer
	nums   []int
	labels []*core.Input
	ch     chan struct{}
	finish bool
}

func NewUIQuicksort(window *core.Window) *UIQuicksort {
	ui := &UIQuicksort{}
	ui.Widget = core.NewWnd(window, 0, 0, window.Width, window.Height)
	ui.nums = []int{12, 4, 5, 9, 1, 6, 8, 11, 3, 7}
	ui.labels = make([]*core.Input, 0)
	w := (window.Width - 6*len(ui.nums)) / len(ui.nums)
	h := window.Height / 20
	for i, _ := range ui.nums {
		label := &core.Input{Wnd: core.Wnd{X: 3 + i*(w+5), Y: window.Height - h*ui.nums[i] - 40, Width: w, Height: h * ui.nums[i], Text: strconv.Itoa(ui.nums[i])}}
		ui.labels = append(ui.labels, label)
		ui.Widget.AddChild(label)
	}
	ui.Widget.Update()
	go ui.quicksort(0, ui.nums)
	ui.ch = make(chan struct{})
	ui.Timer = &core.Timer{TimeOut: ui.Timeout}
	ui.Timer.Start(1000)
	return ui
}

func (ui *UIQuicksort) Timeout() {
	if ui.finish {
		close(ui.ch)
		fmt.Println("finish sort")
		core.MessageBoxInfo(ui.Widget, "info", "finish sort")
		return
	}
	ui.UpdateNums()
	ui.ch <- struct{}{}
}

func (ui *UIQuicksort) UpdateNums() {
	w := (ui.Widget.Width - 6*len(ui.nums)) / len(ui.nums)
	h := ui.Widget.Height / 20
	for i, _ := range ui.nums {
		ui.labels[i].SetText(strconv.Itoa(ui.nums[i]))
		ui.labels[i].SetPos(3+i*(w+5), ui.Widget.Height-h*ui.nums[i]-40, w, h*ui.nums[i])
	}
	ui.Widget.Update()
}

func (ui *UIQuicksort) quicksort(index int, left []int) {
	if len(left) <= 1 {
		return
	}
	i := 1
	j := len(left) - 1
	key := left[index]

	fi, lj := -1, -1
	for {
		if i >= j {
			if left[i] < key {
				left[i], left[index] = left[index], left[i]
			}
			break
		}

		if left[i] >= key {
			fi = i
		} else {
			i++
		}
		if left[j] <= key {
			lj = j
		} else {
			j--
		}
		if fi > -1 && lj > -1 {
			left[fi], left[lj] = left[lj], left[fi]
			fi = -1
			lj = -1
		}
	}
	<-ui.ch
	if i < len(left) {
		ui.quicksort(0, left[0:i])
		ui.quicksort(0, left[i:len(left)])
	}

	if len(left) == len(ui.nums) {
		ui.finish = true
		ui.Timer.Stop()
	}
}
