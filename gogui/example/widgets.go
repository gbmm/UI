package example

import (
	"UI/gogui/core"
	"fmt"
	"math/rand"
	"strconv"
)

type UI struct {
	Widget      *core.Wnd
	Btn1        *core.Button
	Btn2        *core.Button
	Btn3        *core.Button
	Label       *core.Label
	Timer       *core.Timer
	Input       *core.Input
	Combox      *core.Combox
	Checkbox    *core.CheckBox
	Progressbar *core.ProgressBar
	bv1         bool
	index       int
}

func NewUIWidgets(window *core.Window) *UI {
	ui := &UI{}
	ui.Widget = core.NewWnd(window, 0, 0, 640, 480)
	ui.Btn1 = &core.Button{Wnd: core.Wnd{X: 100, Y: 100, Width: 100, Height: 50, Text: "BTN1"}, Click: ui.Click1}
	ui.Btn2 = &core.Button{Wnd: core.Wnd{X: 100, Y: 200, Width: 100, Height: 50, Text: "BTN2"}, Click: ui.Click2}
	ui.Btn3 = &core.Button{Wnd: core.Wnd{X: 100, Y: 300, Width: 100, Height: 50, Text: "BTN3"}, Click: ui.Click3}
	ui.Label = &core.Label{Wnd: core.Wnd{X: 300, Y: 200, Width: 100, Height: 50, Text: "label"}}
	ui.Input = &core.Input{Wnd: core.Wnd{X: 300, Y: 100, Width: 200, Height: 30}}
	ui.Combox = &core.Combox{Wnd: core.Wnd{X: 300, Y: 30, Width: 200, Height: 30}, Click: ui.ComboxClick}
	ui.Checkbox = &core.CheckBox{Wnd: core.Wnd{X: 300, Y: 150, Width: 100, Height: 30, Text: "Check"}, Checked: ui.Checked}
	ui.Progressbar = &core.ProgressBar{Wnd: core.Wnd{X: 50, Y: 400, Width: 500, Height: 30}, Finish: ui.BarFinish}
	ui.Timer = &core.Timer{TimeOut: ui.TimeOut}
	ui.Widget.AddChild(ui.Btn1)
	ui.Widget.AddChild(ui.Btn2)
	ui.Widget.AddChild(ui.Btn3)
	ui.Widget.AddChild(ui.Label)
	ui.Widget.AddChild(ui.Input)
	ui.Widget.AddChild(ui.Combox)
	ui.Widget.AddChild(ui.Checkbox)
	ui.Widget.AddChild(ui.Progressbar)
	ui.Widget.Update()
	ui.Timer.Start(100)
	return ui
}

func (ui *UI) Click1() {
	ui.Btn2.SetVisible(ui.bv1)
	fmt.Println("btn1 clicked", ui.Btn2.Hidden)
	ui.bv1 = !ui.bv1
}

func (ui *UI) Checked(f bool) {
	fmt.Println(f)
}

func (ui *UI) ComboxClick(text string) {
	ui.Label.SetText(text)
	fmt.Println("comox ", text)
}

func (ui *UI) Click2() {
	fmt.Println("btn2 clicked", ui.Input.Focus)
	core.MessageBoxInfo(ui.Widget, "Title", "this is msg box")
}

func (ui *UI) Click3() {
	num := rand.Intn(10)
	ui.Label.SetText(strconv.Itoa(num))
	ui.Input.SetText(strconv.Itoa(num))
	ui.Combox.AddItem(strconv.Itoa(num))
	fmt.Println("btn3 clicked", ui.Label.Text, ui.Label.Hidden)
}

func (ui *UI) TimeOut() {
	num := rand.Intn(10)
	ui.Label.SetText(strconv.Itoa(num))
	ui.Progressbar.SetProgress(ui.index)
	ui.index += 1
}

func (ui *UI) BarFinish() {
	ui.Timer.Stop()
}

func updatePix(window *core.Window) {
	//widget := core.NewWnd(window, 0, 0, 640, 480)
	//btn := core.Button{Wnd: core.Wnd{X: 100, Y: 100, Width: 100, Height: 50, Text: "BTN1"}, Click: Click1}
	//btn2 := core.Button{Wnd: core.Wnd{X: 100, Y: 200, Width: 100, Height: 50, Text: "ABC2"}, Click: Click2}
	//widget.AddChild(&btn)
	//widget.AddChild(&btn2)
	//widget.Update()
	// btn2 := core.Button{Wnd: widget}

	//Wnd: core.Wnd{Phy: &display, X: 100, Y: 100, Width: 100, Height: 50}
	//display.DrawHLine(1, 620, 479, core.Color{R: 255})
	//display.DrawHLine(1, 620, 1, core.Color{R: 255})
	//display.DrawVLine(1, 480, 620, core.Color{G: 255})
	//display.DrawVLine(1, 480, 1, core.Color{G: 255})

	// btn.SetPos(100, 100, 100, 50)
	// btn2.SetPos(100, 200, 100, 50)

	// btn2.Update()
	//
	// display.FillRect(100, 100, 150, 150, core.Color{G: 200})
	//
	// display.FillRect(200, 200, 100, 100, core.Color{G: 200})
}
