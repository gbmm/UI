package example

import (
	"UI/gogui/core"
	"math/rand"
	"strconv"
	"time"
)

type Weather struct {
	Widget *core.Wnd
	Label  *core.Label
	Label2 *core.Label
	Btn    *core.Button
	Input  *core.Input
}

func NewUIWeather(window *core.Window) *Weather {
	ui := &Weather{}
	ui.Widget = core.NewWnd(window, 0, 0, window.Width, window.Height)
	ui.Label = &core.Label{Wnd: core.Wnd{X: 40, Y: 10, Width: 240, Height: 244}}
	ui.Label2 = &core.Label{Wnd: core.Wnd{X: 290, Y: 10, Width: 240, Height: 244}}
	ui.Input = &core.Input{Wnd: core.Wnd{X: 70, Y: 300, Width: 200, Height: 30}}
	ui.Btn = &core.Button{Wnd: core.Wnd{X: 300, Y: 300, Width: 180, Height: 30, Text: "temperature"}, Click: ui.getTemperature}

	ui.Widget.AddChild(ui.Label)
	ui.Widget.AddChild(ui.Label2)
	ui.Widget.AddChild(ui.Btn)
	ui.Widget.AddChild(ui.Input)
	ui.Label.SetImage("E:\\go\\src\\UI\\gogui\\example\\image\\w3.bmp")
	ui.Label2.SetImage("E:\\go\\src\\UI\\gogui\\example\\image\\w5.bmp")
	ui.Widget.Update()
	return ui
}

func (ui *Weather) getTemperature() {
	rand.Seed(time.Now().Unix())
	temp := rand.Intn(100)
	ui.Input.SetText(strconv.Itoa(temp))
	path := "D:\\awesomeProject\\src\\UI\\gogui\\example\\image\\w"
	ui.Label.SetImage(path + strconv.Itoa(temp%7) + ".bmp")
	ui.Label2.SetImage(path + strconv.Itoa(temp%7+7) + ".bmp")
}
