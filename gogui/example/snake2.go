package example

import (
	"UI/gogui/core"
	"fmt"
	"math/rand"
	"strconv"
)

/*使用 label/input 实现*/

type Point2 struct {
	X int
	Y int
}

type UISnake2 struct {
	Widget       *core.Wnd
	Label        *core.Input
	Timer        *core.Timer
	Point2       *Point2
	bv1          bool
	index        int
	size         int
	blocks       []Point2
	color        core.Color
	orientation  string //
	X            int
	Y            int
	flagGameover bool
	labels       []*core.Input
}

func NewUISnake2(window *core.Window) *UISnake2 {
	ui := &UISnake2{}
	ui.Point2 = &Point2{}
	ui.size = 20
	ui.X = 28
	ui.Y = 21
	ui.Point2.X = rand.Intn(ui.X)
	ui.Point2.Y = rand.Intn(ui.Y)
	ui.orientation = "d"
	ui.blocks = make([]Point2, 0)
	ui.labels = make([]*core.Input, 0)
	ui.color = core.Color{R: 240, G: 240, B: 240}
	ui.blocks = append(ui.blocks, Point2{11, 10})
	ui.blocks = append(ui.blocks, Point2{10, 10})
	ui.Widget = core.NewWnd(window, 0, 0, window.Width, window.Height)
	ui.Label = &core.Input{Wnd: core.Wnd{X: ui.Point2.X*ui.size + 3, Y: ui.Point2.Y*ui.size + 3, Width: ui.size, Height: ui.size}}
	ui.DrawSnake()
	ui.Widget.CustomKeyEvent = ui.KeyEvent
	ui.Widget.AddChild(ui.Label)
	ui.Widget.Update()
	ui.Timer = &core.Timer{TimeOut: ui.Timeout}
	ui.Timer.Start(500)
	return ui
}

func (ui *UISnake2) Timeout() {
	ui.SnakeMove()
	ui.Widget.Update()
}

func (ui *UISnake2) DrawApple() {
	ui.Label.SetPos(ui.Point2.X*ui.size+3, ui.Point2.Y*ui.size+3, ui.size, ui.size)
}
func (ui *UISnake2) DrawSnake() {
	num := len(ui.blocks) - len(ui.labels)
	for i := 0; i < num; i++ {
		label := &core.Input{Wnd: core.Wnd{X: 30, Y: 30, Width: ui.size, Height: ui.size}}
		ui.labels = append(ui.labels, label)
		ui.Widget.AddChild(label)
	}
	for j, p := range ui.blocks {
		ui.labels[j].SetPos(p.X*ui.size+3, p.Y*ui.size+3, ui.size, ui.size)
	}
	ui.Widget.Update()
}

func (ui *UISnake2) KeyEvent(key int) {
	if ui.flagGameover {
		return
	}
	ori := string(rune(key))
	nowOri := ori
	if (ori == "a" && ui.orientation == "d") || (ori == "d" && ui.orientation == "a") ||
		(ori == "w" && ui.orientation == "s") || (ori == "s" && ui.orientation == "w") {
		nowOri = ui.orientation
	}
	if nowOri == ui.orientation {
		ui.SnakeMove()
	}
	ui.orientation = nowOri
	ui.Widget.Update()
}

func (ui *UISnake2) SnakeMove() {
	dx, dy := 0, 0
	switch ui.orientation {
	case "d":
		dx = dx + 1
	case "a":
		dx = dx - 1
	case "w":
		dy = dy - 1
	case "s":
		dy = dy + 1
	}
	p := ui.blocks[0]
	if p.X+dx >= ui.X || p.X+dx < 0 || p.Y+dy < 0 || p.Y+dy >= ui.Y {
		ui.flagGameover = true
		ui.Timer.Stop()
		fmt.Println("game over")
		core.MessageBoxInfo(ui.Widget, "info", "game over!score:"+strconv.Itoa(len(ui.blocks)))
		return
	}
	p.X += dx
	p.Y += dy
	if p.X == ui.Point2.X && p.Y == ui.Point2.Y {
		ui.blocks = append(ui.blocks, p)
		for i := len(ui.blocks) - 1; i > 0; i-- {
			ui.blocks[i] = ui.blocks[i-1]
		}
		ui.blocks[0] = p
		ui.Point2.X = rand.Intn(ui.X)
		ui.Point2.Y = rand.Intn(ui.Y)
		ui.DrawApple()
		ui.SnakeMove()
		interval := 500 - len(ui.blocks)*10
		if interval < 60 {
			interval = 60
		}
		ui.Timer.Restart(interval)
	} else {
		for i := len(ui.blocks) - 1; i > 0; i-- {
			ui.blocks[i] = ui.blocks[i-1]
		}
		ui.blocks[0] = p
	}
	ui.DrawSnake()
}
