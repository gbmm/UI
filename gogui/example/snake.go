package example

import (
	"UI/gogui/core"
	"fmt"
	"math/rand"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type UISnake struct {
	Widget       *core.Wnd
	Label        *core.Label
	Timer        *core.Timer
	point        *Point
	bv1          bool
	index        int
	size         int
	blocks       []Point
	color        core.Color
	orientation  string //
	X            int
	Y            int
	flagGameover bool
}

func NewUISnake(window *core.Window) *UISnake {
	ui := &UISnake{}
	ui.point = &Point{}
	ui.size = 20
	ui.X = 28
	ui.Y = 21
	ui.point.X = rand.Intn(ui.X)
	ui.point.Y = rand.Intn(ui.Y)
	ui.orientation = "d"
	ui.blocks = make([]Point, 0)
	ui.color = core.Color{R: 240, G: 240, B: 240}
	ui.blocks = append(ui.blocks, Point{10, 10})
	ui.blocks = append(ui.blocks, Point{11, 10})
	ui.Widget = core.NewWnd(window, 0, 0, window.Width, window.Height)
	ui.Label = &core.Label{Wnd: core.Wnd{X: 3, Y: 3, Width: 570, Height: 430}}
	ui.Label.DrawEvent = ui.LabelDrawEvent
	ui.Widget.CustomKeyEvent = ui.KeyEvent
	ui.Widget.AddChild(ui.Label)
	ui.Widget.Update()
	ui.Timer = &core.Timer{TimeOut: ui.Timeout}
	ui.Timer.Start(500)
	return ui
}

func (ui *UISnake) LabelDrawEvent() {
	//ui.Label.DrawRect(ui.Label.X, ui.Label.Y, ui.Label.Width, ui.Label.Height, core.Color{})
	ui.DrawApple()
	ui.DrawSnake()
}

func (ui *UISnake) Timeout() {
	ui.SnakeMove()
	ui.Widget.Update()
}

func (ui *UISnake) DrawApple() {
	ui.Label.FillRect(ui.point.X*ui.size+3, ui.point.Y*ui.size+3,
		ui.point.X*ui.size+23, ui.point.Y*ui.size+23, ui.color)
}
func (ui *UISnake) DrawSnake() {
	for _, p := range ui.blocks {
		ui.Label.FillRect(p.X*ui.size+3, p.Y*ui.size+3,
			p.X*ui.size+23, p.Y*ui.size+23, ui.color)
	}
}

func (ui *UISnake) KeyEvent(key int) {
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

func (ui *UISnake) SnakeMove() {
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
		//core.MessageBoxInfo(ui.Widget, "info", "game over!")
		ui.Label.SetText("game over!!!  score:" + strconv.Itoa(len(ui.blocks)))
		return
	}
	p.X += dx
	p.Y += dy
	if p.X == ui.point.X && p.Y == ui.point.Y {
		ui.blocks = append(ui.blocks, p)
		for i := len(ui.blocks) - 1; i > 0; i-- {
			ui.blocks[i] = ui.blocks[i-1]
		}
		ui.blocks[0] = p
		ui.point.X = rand.Intn(ui.X)
		ui.point.Y = rand.Intn(ui.Y)
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

	// ui.Widget.Update()
}
